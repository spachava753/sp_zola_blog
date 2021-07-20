+++
title =  "A Gopher's foray into Rust"
date = 2021-06-29
+++

# Points

- [x] Why I decided to look into rust now
- [ ] What I tried to do
    - [ ] partial port of a private project
    - [ ] results of port
- [ ] What I liked
- [ ] What I disliked
- [ ] How I learned (separate post later)

---

You may have heard of Rust. If not, here's the slogan of the language (as of 2021):

> A language empowering everyone to build reliable and efficient software.

It's quite an ambitious statement. However, after using it and understanding what values Rust provides to a programmer,
I was utterly, completely convinced.

# Why now

This is not my first foray into Rust. I've read the Rust book over the past two years a couple of times or so out of
sheer curiosity, as well as passively follow the language's development. I was drawn to the idea of no null or nil. I've
experienced the idea of errors as values in Go, and wanted to see how Rust implements this concept. One particular
interest is that Rust was
the [most loved language](https://insights.stackoverflow.com/survey/2020#technology-most-loved-dreaded-and-wanted-languages-loved)
5 years running as of 2020. Yet, I've never actually taken the dive to write code.

When I first read the Rust book two years ago (early 2019), it quickly spurned my idea of diving
headfirst into a project. The **very** verbose syntax looked difficult to comprehend in my first foray, less so the
second time. After discovering the 2020 Stack Overflow developer survey, and
the [blog post](https://stackoverflow.blog/2020/06/05/why-the-developers-who-use-rust-love-it-so-much/) explaining why
Rust tops the most loved language five years in a row, I decided to take stab trying Rust again. I read the Rust book
once more, which was much less confusing than when I read the first time. Due personal circumstances, I didn't follow up
and just ended my second foray right there.

It wasn't until recently that I made my third foray into Rust. This time, I wanted to commit and implement some kind of project to actually learn the language. I was inspired by the constant stream of blog or reddit posts that reported a multi-fold increase in performance, the reliability of deployed services written Rust, as well as a number of highprofile projects written in Rust (one example is [firecracker](https://github.com/firecracker-microvm/firecracker). The first thing I did, was go through the Rust book...again. This time around, I finished it in about a day, as I was able to grok most of the concepts easily by drawing on previous experiences. The next thing I did was read a number of blog posts and Rust books. While doing so, I was searching for a viable project to start. I wanted something non-trivial, something that takes more than an evening, or a couple of "getting started" tutorials. Finally, I settled on a port of a project that required high performance and reliability.

# The Project

