+++
title = "Microservices 😱"
date = 2021-01-24
[taxonomies]
tags = ["Microservices"]
+++

Writing microservices is hard. Doesn't matter what language, going from monolith to microservices means introducing more moving parts, more possibilities for failure, yet in order to scale effectively, they are a necessary evil. According to this [post](http://highscalability.com/blog/2014/4/8/microservices-not-a-free-lunch.html) on the High Scalability blog, authored by Benjamin Wootton, CTO of [Contino](https://www.contino.io/): the definition microservices is this:
>Microservices are a style of software architecture that involves delivering systems as a set of very small, granular, independent collaborating services

What the heck does that mean? Well, lets break it down.
<Blockquote>Microservices are a style of software architecture...</Blockquote>

Okay, so microservices is way to build software, just like monoliths are another way to build software. In that case, microservices don't refer to any one programming language or platform.

>...that involves delivering systems...

It looks like creating software isn't enough. We also have to figure out how to deliver or make it available to our users. Who woulda thunk?

>...as a set of very small, granular, independent collaborating services

This is the most important part. Microservices, true to its name, means building small, purposeful services that are really meant to a single task. What this task is and the scope of this task is up to the creator of microservices but like Unix philosophy, microservices are meant to do one thing each, and do it well. Of course, each microservices can't achieve much in isolation. as such, they should also collaborate.

Okay. To sum it all up, microservices is a style of building and delivering software that does one thing well and collaborates with other microservices. Sounds good. 👍

Thank you for reading! Be sure come back for the next post...I'm just joking 😅. There is still much more left to think about. 

Lets start with the "collaboration" part. How are your microservices supposed to communicate? Based on your system, perhaps a message bus or a pub sub makes more sense. Maybe you fancy GRPC, or you just want good all REST. Heck, it could be a conglomeration of all of these! So we decided on a means of communication between our microservices. Are we done?

Wait! What about observability? When we had monolithic architecture, we knew where all requests go the same place, but now that our code is split into separate units, we can't keep track of where each request ends up and what services it uses! What if some requests are failing? How do we know what service caused to request to fail? Often, a single microservices might talk to not just one, but *multiple* services, so it can get pretty difficult trace where each request goes.

Lets assume that we have taken care of observability, now its time to talk about monitoring. What's the difference between observability and monitoring? Observability allows to obtain insights into how requests flow through our system. Monitoring allows us to obtain metrics about how each of our services are performing, metrics like: RPS, CPU usage, Memory usage, Number of Bad requests, Average duration, and much more. Basically, if it is possible to monitor it, its not bad idea to monitor it 🤷‍♂️.

We're still missing something...oh yeah! Scalability! One of the most important reasons to go through all this trouble in the first place. If we used a monolithic architecture, scaling that service usually meant acquiring bigger specs like more memory, more cpu cores, and more network bandwidth. However, there is a limit to this, and we get diminishing returns as we keep upgrading. However, if we split our service into microservices, we scale individual microservices to handle load needed. However, we might still hit a limit for one or more microservices, after all, the problem of getting more and more expensive hardware hasn't gone away. What a troubling issue! 

Hm 🤔🤔, what if...nah...maybe...hear me out. Lets run... multiple instances of microservices. I know, I know, it sounds insane, but stay with me. If we were able to run multiple instances, or copies of a single microservice, this would allow to simple acquire more servers instead of dishing out cash for high-spec hardware. As Google calls it, we could run on commodity hardware, that is, buy a couple of used desktops off Ebay, connect them together, and run our microservices on them. Wonderful 🤗! But wait, theres more! How do the other services know where to find our running services? After all, there are multiple copies running at different locations. This is a well known problem called service discovery. But how do we solve this problem?

Besides these concerns, we also have others like encrypting connections between microservices for security, figuring out how to deploy our microservices after we create a new version, setting up a central location for all of our microservices configuration, load-balancing and much more. This whole microservices thing is really starting to daunting...

As it turns out, people realized that this whole microservices thing can only work if all of these concerns were addressed. One solution to this problem is Netflix's OSS stack. It is Java-based, and solves the problems mentioned above, like service discovery, central configuration, monitoring, etc. It's most often used with Spring Cloud, which is a set of libraries built to work with the (in?)famous Spring framework. We still haven't addressed one issue: deployment! How are ever supposed to actually deploy our microservices to our Ebay servers?! I guess we can use ssh or ansible or something 😦.

We haven't even got to the part about increasing and decreasing instance of microservice based on demand...😱

# Fast forward to 2020 ⏩

Its the age of the cloud! We have AWS from Amazon, GCP from Google, Azure from Microsoft, IBM Cloud from IBM, Alibaba Cloud from Alibaba, etc. The cloud has enabled companies big and small to get high quality servers in seconds, while getting charged for only the amount of time you use a server! Even crazier, Google came out with a project called Kubernetes, which really just took the (microservices) world by storm. Microservices are now packaged into units of deployment called containers, which contain all of the necessary software that your application needs, so you run it anywhere, regardless of the deployed servers software and library versions. Kubernetes takes these units of deployment, and "orchestrates" them. What does this mean? It takes care of scheduling, deploying, and running your microservice on a server, which you don't even need to specify, Kubernetes just figures it out! How crazy is that?! No, here's the crazy part: you remember all of those annoying problems -- I mean concerns that we saw earlier? Kubernetes takes care of all of that! Yes, not only does it schedule and run your microservice, it takes care of service discovery, configuration, scaling on demand, and much more!🤯  ...Sort of. It's true that Kubernetes can take care of all our problems, but we need to install some software that helps Kubernetes help us. An example of this is a service mesh. In essence, service meshes do the things that Kubernetes doesn't do, things like encrypting communications between microservices, circuit breaking, monitoring, observability, etc. Kubernetes, like microservices, was made to do one thing: manage the lifecycle of containers. Hence, the need for software like service meshes.

One small problem: setting up Kubernetes for production can be quite complicated and tedious. This is where the cloud can step in. The cloud can help set up Kubernetes for us, so that it is resilient against network outages and database failures. Our choice of cloud can also add or remove servers under Kubernetes automatically, so we can run only as many servers as needed. Great! Previously, we might have relied on something like the Netflix OSS stack, which often had constraints like first class support for a single language or required to add something to our source code that isn't directly related to the logic contained in our microservice.