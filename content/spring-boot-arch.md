+++
title = "Spring Boot Architecture"
date = 2018-12-19
[taxonomies]
tags = ["Spring Boot"]
+++

## A Bootiful Introduction

(pun intended)

Although Java is considered notorious for being a verbose language, very few languages come close to the sheer size of the community and features that Java has to offer. One of the best things that the Java community and ecosystem has to offer is frameworks. If you are a Java developer, you probably heard of Spring Boot. Spring Boot is a wonderful framework that takes the pain out of writing boilerplate code for web apps. The framework uses an opinionated, convention-over-configuration development approach that makes it blazing fast to develop on projects. Take any enterprise back-end, or a microservice in 2018, and there is a good chance it is written with Spring Boot. Today we will be looking at the way your application should separate its code into multiple layers.

## Architecture Layers

So what are layers anyway? They are a simple way to separate your code into cohesive parts that have a distinct job. Layers also allow you define your application while still in the designing process, so you don’t have to worry about how to organize the code.

### 3 is the magic number

No matter what type of application you are writing, three layers are the minimum. Naturally, you may have more than three layers based on the complexity of your project. The reason why three layers is the minimum is because each layer corresponds to the part that your application should be handling.

### The Web layer (Presentation layer)

The presentation layer is basically where you write the code to expose your application. The exposure could be through REST API or serving web pages. This layer should only contain code that handles the requests to your app and responses that it gives back. Spring Boot offers an easy way to define controllers that serve template HTML pages or REST API through annotations, making it very easy to expose your application. Some applications may not have a presentation layer as the creator did not want clients interfacing with it. Instead, the application might use other forms of communication such as a messaging system. Keep in mind that a messaging system is not considered part of the presentation layer.

### The Service Layer (Business Logic layer)

The Service layer is where all the magic happens. The Service layer contains all of your business logic, and the any data that your presentation layer needs to persist and retrieve passes through the service layer. The service layer has two components, the service interface and the interface implementation. This allows the Spring Framework to use Dependency Injection to inject your service to components that use it. The separation of the definition and the implementation of the services mean that you can have different implementations for a given service, allowing you to choose during run time what service you want.

### The Repository Layer (Persistence layer)

The Repository layer is self-explanatory: its where you put your code to persist your data. Spring Boot has made astonishingly simple to write code to persist your data by simple writing an interface the extends the JpaRepository interface. Spring Boot makes use of a concept known as ORM: Object Relational Mapper. ORMs translate objects into records that can be stored inside a relational store.

## A Layer for you… and you… and you…

Three layers are a must for any application to start, but sometimes an application can become more complex than expected, creating even more layers. Sometimes you may need an extra layer to translate data that comes in from the presentation layer to send to the service layer. Other times, you may need an extra layer for listening in on messages coming in on a Kafka topic. It all depends on what your application is doing.

## A Reference App

I built a reference app that follows the principle outlined above, and more. The app has auditing, REST API documentation, and custom repository implementations. The app is built on the idea of system that stores user accounts. Check it out at my [Github](https://github.com/spachava753/spring-boot-reference-architecure).
