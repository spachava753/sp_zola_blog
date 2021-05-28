+++
title = "Terraform: The Good and The Bad"
date = 2018-12-06
[taxonomies]
tags = ["Terraform", "Infrastructure Provisioning"]
+++

# Intro

So … what is Terraform? In order to know what Terraform is, we need to know what infrastructure provisioning is. **Infrastructure provisioning** is the process of acquiring infrastructure to run your software in. Before the concept of Infrastructure-as-Code came in, companies had acquire, configure, and maintain their own hardware, and have a dedicated team to do all of these things. Now, with infrastructure provisioning tools, companies can build, version, and configure their hardware all through code. With cloud-native or hybrid-cloud infrastructure exploding with popularity, infrastructure provisioning tools flourished. These tools allowed companies to replace the cumbersome and tedious cycle of ordering new hardware, waiting for it to arrive, configuring it, and then later upgrading it with simply working with some code. Now, all it takes is some code to get a virtual machine on one of the cloud vendors. Some vendors provide infrastructure provisioning tools with their platform, such as AWS with Cloudformation. This forces customers to vendor lock-in, which is not ideal, as this cause problems in the future if you want to expand into other vendors, or perhaps migrate to another platform. Other tools are **cloud-agnostic**, meaning the tool does not care what platform you are on. Terraform is one such tool.

# The Good

## Cloud-agnostic

I would like to re-iterate that while vendor lock-in is acceptable, it is not recommended for established institutions. It’s better to keep your options open, especially if you find out that a different platform could potentially suite you better in the future. That said, Terraform allows to you to provision on a number of “providers”, Terraform’s lexicon for platforms.

## Modules and Variables

A major feature of Terraform is DRY, Don’t-Repeat-Yourself. And Terraform modules truly represent this concept. Modules in Terraform are basically reusable components that can be used to easily create any resources you need, without having to write the same code again and again. Discovering Terraform modules to use is a cinch as well, leaving you with no excuse to repeat your code. A big plus in Terraform is variables. Instead of having to go into each of your files to change a single value that is repeated tens, or hundreds of times, just use a variable, and change the value of that variable.

## Simple. Delightfully simple

Terraform is terrifyingly simple, yet astonishingly powerful. If you compare a .tf file, Terraform’s file extension, and a Cloudformation template, they are world’s apart. In just a couple of lines, you could achieve a lot more than if you were writing a Cloudformation template. Plus, the docs on their website are simple enough for anyone to follow, even inexperienced novices. One of things that Terraform does really well is build your resources based on a graph that depicts all the dependencies of your resources, so all you have to worry about is actually coding your resources, without worrying about what depends on what. Another thing I really appreciate is that its easy to provision hundreds of resources, and at the same time, destroy it as well. Keep that in mind. _All it takes is one command to create and destroy hundreds of resources._

# The Bad

## Simple. Horrifyingly simple

One of the best things about Terraform is that it is really simple, making it approachable and friendly for people who want to learn it. However, if a person does not have sufficient knowledge about the platform they are provisioning on, it can be really easy to mess up. On AWS Cloudformation, there are tons of lines you need to write, requiring prerequisite knowledge about AWS, Cloudformation, and how each resource intertwines with another. This acts as a barrier when provisioning resources for your enterprise as it will be a significantly harder endeavor for novices to simple finagle their way to the end. Terraform on the other hand, is not so strict, and in fact, will use any defaults it has when not provided with sufficient information. This can pose a major issue as someone who doesn’t have a full handle on what’s going on could still build your infrastructure, ignoring any security concerns and best practices that keep your company safe. Another major issue is the destruction of resources. All it takes is:

    terraform destroy

and now your entire infrastructure is out the window. Of course, this is an extreme case, but I just wanted to stress how easy it is to wantonly destroy and create resources.

# The Reality

While I make it seem as if the bad parts of Terraform can cause a company to go bankrupt, its really not that bad. The truth is, it is such a pleasure working with Terraform, that any piddling disadvantages of using it shrink away in the face of the true potential of Terraform. It’s true that Terraform is simple, and anyone can learn it, but a person can’t fully utilize the potential of Terraform if he or she doesn’t have a full understanding of the platform they are working on. I recommend using Terraform for your infrastructure provisioning needs, as it is one of the best tools out there.
