+++
title =  "A Gopher's foray into Rust"
date = 2021-06-29
+++

# Points

You may have heard of Rust. If not, here's the slogan of the language (as of 2021):

> A language empowering everyone to build reliable and efficient software.

It's quite an ambitious statement. However, after using it and understanding what values Rust provides to a programmer,
I was utterly, completely convinced.

# Why now

This is not my first foray into Rust. I've read the Rust book over the past two years a couple of times or so out of
sheer curiosity, as well as passively follow the language's development. I was drawn to the idea of no null or nil. I've
experienced the idea of errors as values in Go and wanted to see how Rust implements this concept. One particular
interest is that Rust was
the [most loved language](https://insights.stackoverflow.com/survey/2020#technology-most-loved-dreaded-and-wanted-languages-loved)
5 years running as of 2020. Yet, I've never actually taken the time to write code.

When I first read the Rust book two years ago (early 2019), it quickly spurned my idea of diving headfirst into a
project. The **very** verbose syntax looked difficult to comprehend in my first foray, less so the second time. After
discovering the 2020 Stack Overflow developer survey, and
the [blog post](https://stackoverflow.blog/2020/06/05/why-the-developers-who-use-rust-love-it-so-much/) explaining why
Rust tops the most loved language five years in a row, I decided to take a stab trying Rust again. I read the Rust book
once more, which was much less confusing than when I read it the first time. Due to personal circumstances, I didn't follow up
and just ended my second foray right there.

It wasn't until recently that I made my third foray into Rust. This time, I wanted to commit and implement some kind of
project to learn the language. I was inspired by the constant stream of blog or Reddit posts that reported a
multi-fold increase in performance, the reliability of deployed services written Rust, as well as several
high-profile projects written in Rust (one example is [firecracker](https://github.com/firecracker-microvm/firecracker).
First order of business: going through the Rust book...again. This time around, I finished it in about a day, as I was
able to grok most of the concepts easily by drawing on previous experiences. The next thing I did was read a number of
blog posts and Rust books. While doing so, I was searching for a viable project to start. I wanted something
non-trivial, something that takes more than an evening, or a couple of "getting started" tutorials. Finally, I settled
on a port of a project that required high performance and reliability.

# The Project

The project I chose was to port a private application that takes XML data from a URL and massages it into different JSON
representations depending on the information requested. The actual business logic is simple, but the application
requirements are demanding. It needs to serve thousands of requests per second consistently with minimal latency and
handle bursts of traffic. Often the JSON objects returned in each request can be 60 to 70 kilobytes uncompressed. The
original application is a Java Sprint Boot app, with multiple instances of the app running across many servers.
It was then partially ported to Golang to measure the performance gains, and the effort needed to develop new features.
The Golang application showed impressive results, besting the Java application with multi-fold the performance. The
Golang application was able to handle a sustained 30,000 requests per second on my desktop. My desktop has an AMD Ryzen
9 3900x 12-core processor with 24 threads and 32 gigabytes of RAM running Ubuntu 20.04. For the
experiment, I set the limit of open file handles to unlimited and used [vegeta](https://github.com/tsenart/vegeta) as
the load generator.

I started implementing the same feature set that the Golang port implemented, so I could have a fair comparison. Along
the way, I ran into multiple pleasant and not-so-pleasant surprises.

The first thing I did was set up an Actix Web 3 app that served a "hello" endpoint. There was no other special reason
other than the fact the Actix Web has been battled-tested and been deployed in production for quite some time now, and
is a mature project. There are also plenty of examples to learn from. It was fairly simple to set up, but I quickly ran into
issues with the async runtime versions when I tried to contact a MySQL database. I wanted to
use [sqlx](https://github.com/launchbadge/sqlx), which was the only mature *async* SQL project at the time. I also
really liked some features of the project like compile-time checking the SQL queries. Most of all, I liked the fact that
I didn't need to create a separate blocking thread to run [diesel](https://github.com/diesel-rs/diesel) SQL queries.
However, the latest version of sqlx (which was 0.5.x at the time) used the Tokio runtime 1.0, which was incompatible with
the Actix runtime, which relied upon the 0.2.x version of the Tokio runtime. To make sqlx work, I had to use
one version prior, 0.4.x which used the 0.2.x Tokio runtime, making it compatible with the Actix runtime. After
continuing to develop the application, I became more and more annoyed with Actix. Actix has its own data types that
behave similarly to `Arc<T>`, which must be used to pass data to each of the "workers" that the Actix runtime creates
per thread, BUT if the data is already thread-safe, then you have to wrap the data with a _different_ type. I pulled
in [reqwest](https://github.com/seanmonstar/reqwest) to fetch the XML data, but it relied on the Tokio 1.x runtime,
which again meant that I had to search for an old enough version that used the Tokio 0.2.x. At this point, I was pretty
annoyed with Actix Web, so I searched for an alternative, and I found [warp](https://github.com/seanmonstar/warp). A
major point of confusion is why Actix implemented its own runtime when Tokio 0.2.x and above is a **work-stealing
scheduler, just like Golang's goroutine runtime** (see the blog
post ["Making the Tokio scheduler 10x faster"](https://tokio.rs/blog/2019-10-scheduler)). I liked the simplicity of
warp, and I was able to upgrade sqlx and reqwest to use the Tokio 1.x runtime. Sqlx version 0.5, in particular, offered a
better API as I was having problems connecting to the database in 0.4.x, which was resolved in the latest version.

I struggled through the rest of the feature set ([serde](https://github.com/serde-rs/serde) in particular was an
absolute joy to use) and eventually matched the feature set implemented in the Golang app. One point of interest is I
could not find any sufficiently advanced local cache libraries like
Golang's [bigcache](https://github.com/allegro/bigcache), which automatically collects stale entries, but I did
find [cached](https://github.com/jaemk/cached). Cached offered a TTL-based read-write lock hashmap, which was good
enough to move forward with the application. The app needs a local cache so the database isn't overwhelmed and the
Golang app used bigcache, so I was somewhat disappointed to find that there weren't any equivalents in the Rust
ecosystem.

After fighting with the compiler, the Rust app reached the same point as the Golang app.
After benchmarking the Rust application it achieved a measly 10,000 requests per second! It also had much, much higher
tail latencies past 2000 requests per second than the Golang app. I was utterly confused at the terrible performance in
comparison to the Golang app and thought it was because I was not using Actix Web which tops
the [Techempower benchmarks](https://www.techempower.com/benchmarks/). For reference, the Golang app used the
famous [fast-http](https://github.com/valyala/fasthttp) package, which was one of the top frameworks in the Techempower
benchmarks in comparison to warp. So I fought with the compiler and incompatible dependencies for a couple of late
nights to convert the Rust to use Actix Web, but the performance was roughly the same. At this point, I was lost as to
why my Rust application was doing so bad until I realized after a week-long break from the project that it was probably
mutex contention. After all, bigcache used a **sharded** hashmap, so I pulled
in [dashmap](https://github.com/xacrimon/dashmap). I went back to using warp and Tokio 1.x, converted the cache to use
dashmap, and the performance shot up to 35,000 requests per second, with half of the latency of the Golang
application **and** very impressive small tail latencies! At this point, I was curious as to what else I could improve "
just" by pulling different dependencies. After searching around, I realized that I could use a different hashing
function, such as [aHash](https://github.com/tkaitchuck/aHash) which is designed to be used "in in-memory hashmaps".
Just by replacing the hash function, the requests per second shot up _again_ to 45,000 with slightly better latencies.

After this point, I stopped as I was burnt out, but there were still options left to explore. For example, at this point in time, dashmap did not use the [parking lot](https://github.com/Amanieu/parking_lot) library, which boasts
significantly better performance than the standard library mutexes used in dashmap:

> When tested on x86_64 Linux, parking_lot::Mutex was found to be 1.5x faster than std::sync::Mutex when uncontended, and up to 5x faster when contended from multiple threads. The numbers for RwLock vary depending on the number of reader and writer threads, but are almost always faster than the standard library RwLock, and even up to 50x faster in some cases

I also didn't bother going back and fixing some particularly gnarly code in the hot path that left some performance gains on the table. 

# Thoughts

## Performance comparison

I will be the first to admit that the performance comparison might not be fair, as I spent significantly longer on the Rust
port than the time spent on Golang port, in part due to the steeper learning curve of Rust's compiler. I've also fiddle around with aspects of the code that was not looked at for the Golang port, such as the hash function for the bigcache. However, this was a learning experiment done out of personal curiosity, and my goal at the start was to see if I can simply match the performance of the Golang port. Having achieved the goal, I just wanted to write up how I got there.  

## Type System: Async, Enumerations, and Null

There were certain points that I liked about the experience that I want to share. First, the rules of the ownership system in Rust is pretty intuitive, at least for an experienced programmer, especially one that has experienced some pitfalls of
concurrency. However, coming from Golang's succinct syntax, I felt less productive when writing Rust code. I am sure it will get better with time and practice, but it is a stark contrast to Go's learning experience where the "Tour of Go" is enough to write decent code. Fracturing in the Rust async ecosystem also made me fatigued, although that's getting much better and can only improve as time goes on. One thing in particular that I found absolutely amazing is the advanced type system. I didn't realize what I was missing out on until I used Rust's enumeration types. Especially with `Option`, which completely removes the [null](https://www.infoq.com/presentations/Null-References-The-Billion-Dollar-Mistake-Tony-Hoare/) problem. With `Option`, you aren't left wondering whether you should bother adding a null (or nil in Go's case) check as a preamble to a function's body whenever you accept references as function arguments, simply because _this is not possible in Rust!_ (You can do this in unsafe Rust, but unsafe Rust code is pretty rare).

After getting some experience in writing async code, I think Rust's way of using `await` syntax made more sense to me. The concept of colored functions  

## Packages

Cargo was an absolute joy to use, and I became jealous of Javascript developers when I realized this luxury is something they have on the daily. I like Go modules, but it took the community _way_ too long to introduce Go modules. Not to mention, I still have to search up the _full path of the library I want to use!_ For example, if I want to use "github.com/pkg/errors", I have to run `go get github.com/pkg/errors` or type out "github.com/pkg/errors" in `go.mod`. The path github.com/pkg/errors is pretty easy to remember, but you can imagine that if the path has more characters with a less memorable name, you would be forced to go to Github or pkg.go.dev, search up the full path of the library, _then_ add it to your project. In contrast to Rust, you would just type the name of the library and the version, like this `serde = 1.0`. This is made easier with cargo's extensibility using plugins. You can install the `cargo-add` plugin, then all you have to do _is run `cargo add serde`_. This may seem like a small improvement, but when working on a non-trivial project, you often have multiple dependencies, and constantly looking up the path can get annoying real fast.

I also like that there is an opinionated way to structure projects. You just run `cargo new` and it creates a new binary or library project, complete with the `src` folder. No need to think about whether you need a `pkg` folder or not.

I simply cannot state how absolutely amazing the cargo tool is. It's the small things that make the developer experience great. Kudos to the Rust community for creating is an amazing tool! 

# Conclusion

My first non-trivial project with Rust was a mixed bag. On one hand, I really liked Rust's expressive type system, as
well as cargo.