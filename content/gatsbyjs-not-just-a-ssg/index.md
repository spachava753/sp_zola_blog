+++
title = "GatsbyJS: Not just a SSG"
date = 2019-08-29
[taxonomies]
tags = ["SSG", "Tutorial", "AppSync", "GraphQl"]
+++
By now, I'm sure everybody has heard of GatsbyJS. It's a React-based web framework with a plugin system that allows you to glue various data sources to page views. Of course, it functions as a SSG (**S**tatic **S**ite **G**enerator), but it's much more powerful than that.

# Another SSG? Really? 

I know what you are thinking. "There are already so many other SSGs out there, what makes Gatsby so special?" One of my favorite things about the framework is the plugin system. You will pretty much find whatever you need through the form of plugins. Another big plus is that it is based on React. It will pre-render the components you define to create a static site. However, since it is based on React, it will also preserve the code you write for run-time, such as code for fetching data in **componentDidMount**. Just these features alone would want to make anyone switch, but I'm sure some might think as a result, Gatsby would be very complex to develop in. **Wrong again!** The Gatsby team created some amazing documentation, and you can get started creating some sample projects. One of the sample projects is the blog you are reading right now! As a developer, you don't have to worry about a complex frontend techstack or a lack of community and support.

# Key features

In an era of where users expect websites to load insanely fast, GatsbyJS rose to the standard to support frontend web optimization and blazing fast performance **out of the box**. As a developer, we don't have to worry about code splitting, service workers / progress web app capabilities, asset optimization, etc. Gatsby automatically takes care of all these things so that you can focus and developing your application. While developing, Gatsby will incrementally build and live-reload, so you can see changes immediately. All of these are part of the "developer experience", as Gatsby has toolset and an active ecosystem. However, if you saw the title, then you will realize that I didn't just create this post to talk about Gatsby's capabilities as a SSG.

# What do you mean, more than a SSG?

According to [gatsbyjs.org](https://www.gatsbyjs.org/):

> Gatsby is a free and open source framework based on React that helps developers build blazing fast websites and apps

Notice that it says that Gatsby is a framework based on React. This means you can use the same flow that you use to develop React applications. There are only a couple of lines of code that you have to add to make Gatsby work, but you don't need to change the way you write code. Frontend developers will rejoice over this. I also mentioned that Gatsby will only pre-render what it can at build time, so you can still create a fetch dynamic data at run-time and populate that then. What this allows you to do is first show the user some kind of feedback, like a pre-rendered header and footer, and dynamically load data. Something like this:

![GIF showing dynamically loaded content](dynamic-data-fetching.gif "GIF showing dynamically loaded content")

We are essentially creeping into the territory of web applications. You would be hard-pressed to be able to do this in other static site generators, especially ones like Hugo that use HTML templates to create static sites.

# So what's the catch?

Like all things in software, tradeoffs must be made. The biggest thing you will be sacrificing is server-side rendering fast changing content. The reason I say "fast-changing content" and not dynamic data is because Gatsby can handle dynamic data to an extent, whether you are fetching the during runtime or during build time. Imagine you run a news platform, and each time someone publishes something, you have to rebuild the site. Authors can publish articles daily, hourly, or even faster than that. GatsbyJs is not built to handle fast-changing content. On the other hand, I would suggest looking at something like Next Js for server-side rendering fast-changing data. For other use cases like an ecommerce site or a portfolio site, I would suggest GatsbyJs as it offers unparalleled flexibility and ease of development like no other framework.
