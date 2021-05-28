+++
title = "Why Kubernetes?"
date = 2020-01-16
[taxonomies]
tags = ["Kubernetes"]
+++

In order to understand why a container orchestration platform is needed, one must first understand why containers are needed. 

# Why Containers?

During the recent years of application development, companies have been heavily invested in containerization and automation. There are many reasons for this movement, but to name a few: fast development speed, highly available applications, and isolation with minimal overhead. Containers were especially popularized as they isolated processes through the use of Linux namespaces, while their usage could be monitored and throttled with cgroups. However, to actually run processes inside namespaces was clunky and unruly. Docker popularized this technology by making it easy to run anything inside containers. Developers no longer have to excuse of "it works on my machine". Containers run in isolation, so they cannot affect other processes running in the same machine. At the same time, they offer portability, so the the same container image can run anywhere. Because of portability and ease of use with container runtimes, introducing automation into the development flow with containers was the next natural step. With all of these benefits, developers and companies quickly caught on and started running their applications in containers. However, in order to run containers in production, especially at a large volume, there has to be a way to orchestrate them. A lot of companies ran into this problem, and one such company, who was an early adopter of Linux namespaces, even before Docker, was Google. They needed to manage a large number of containers at scale, with near infinite horizontal scalability. Their initial experiments quickly led to Borg, the original container orchestration system. After continuous refinements, the successor was Omega. Finally, Omega was open-sourced to become Kubernetes. 

# Why Kubernetes of all things?

Of course, one may ask why Kubernetes is the de facto standard, as there are other alternatives out there. Docker Swarm, Apache Mesos, and even just an in-house solution are all ways to orchestrate containers. However, none of them can quite measure up to Kubernetes for a couple of reasons.

## Scalability

While Docker Swarm could be used to orchestrate containers, it was not ideal for a large volume of containers. It did not make it easy to orchestrate containers, it only offered a solution that might be better than what an in-house solution could do. Apache Mesos is also quite capable, but relies on the concept of providers and frameworks, which means the Mesos does not offer orchestration functionality out of the box. While Kubernetes knows its role of orchestrating containers and centers its entire platform around it, Mesos it built to be more generalized. Kubernetes was built to manage hundreds of thousands of containers comfortably.

## Battle-Tested

The single biggest reason is that Kubernetes is battle-tested, and battle-tested at Google, a company that operates at a global scale with multiple critical systems and 24/7 traffic. Google has already addressed the problems of container orchestration and have been hashing them out for over a decade. Kubernetes provides a sensible, opinionated view to container orchestration with a batteries included approach. Ever since Kubernetes has been open-sourced, it has been used to run multiple production environments and thousands of developers have contributed to its development. There is no better time to pick up Kubernetes than now.

## Elasticity

Another major contributor to Kubernetes is the fact that many major companies needed a way to manage their own data centers, and needed an easy way to deploy applications to those datacenters. The cloud allowed many companies to scale and handle large bursts of traffic without costing an arm and a leg. While this is true for small companies and startups, major enterprises actually prefer to use their own data centers as the primary way to serve traffic while relying on the cloud for bursts of traffic that their data centers cannot handle. A prime example is Dropbox, who initially used AWS S3 to store their data, but later built their own data center called Magic Pocket. Kubernetes allows companies like Dropbox to manage their data centers with ease. Kubernetes can also intelligently manage applications, scaling up certain apps that have more traffic while scaling down other apps that might be idle or handle less traffic. Besides scaling applications, Kubernetes can also dip into the cloud to add workers to the cluster on demand when there is traffic that the cluster cannot handle.

## Extensibility and Customizability

Besides being battle tested, extensibility and customizability is probably the next biggest factor contributing to Kubernetes explosive rise in popularity. While Kubernetes comes with sensible defaults, what it really is at its core is a set of intricately connected policies, protocols and components that manage containers together. Each component in Kubernetes can be modified, or replaced with a custom implementation so long as it conforms to protocols and policies. With CRDS, Custom Resource Definitions, Kubernetes can manage resources that are not shipped with the Kubernetes project. For example, a Website could be a CRD, which will accept a location to a SPA, and will automatically deploy the SPA to an Nginx Container while exposing the container as a service and setting up auto scaling.

## Thriving Community

There are plenty of resources online, from books to courses, that allows a person to go from zero to hero in Kubernetes. Since it is quite a complex system (which makes sense, considering all of its capabilities) it takes time and effort to truly master Kubernetes. However, all of the resources are within reach, and if you are stuck, chances are, some else was stuck in the same place, posted their question to Stack Overflow or a forum, and somebody already answered. There is no lack of documentation and articles on designing and administering clusters.

# So Kubernetes all the way ... right?

While Kubernetes offers a sane solution to manage a large volume of containers, it is not always the best direction for a company to take, especially for small startups. For behemoths and large enterprises, Kubernetes is the best solution and direction to head toward, hands down. But for small startups, it might actually be cheaper and easier to leverage the cloud's services, such as AWS Lambda, which is AWS's serverless offering. Especially with the free tier offered in each of the clouds, startups can run their services with little to no cost. The cloud also offers solutions for different functionalities like batch processing, search, and machine learning workflows which could be much more performant than can what can be developed in-house with Kubernetes.

Overall, anytime it is necessary to orchestrate containers, the de facto choice should be to choose Kubernetes.
