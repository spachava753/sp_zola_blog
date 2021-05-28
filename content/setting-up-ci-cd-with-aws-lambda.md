+++
title = "Setting up CI/CD with AWS Lambda"
date = 2019-08-10
[taxonomies]
tags = ["AWS", "Tutorial", "CI/CD", "Serverless"]
+++
# Intro

I'm sure everybody has heard of AWS Lambda. It's AWS's version of a serverless service. Serverless is gaining traction in the industry from heavyweights to startups. It's ease of creating and deploying functions make it popular to be used whenever possible. I won't go too much into why you should use serverless as there plenty of articles and blog posts dedicated to why it's great. However, as soon as you introduce automation to any technology, trying to manage it becomes 10 times harder. This article is about setting up a CI/CD process with AWS CodePipeline for AWS Lambda. **I created this article to expand on the already comprehensive guide on the AWS docs and to document my experience. The link for the official docs can be found** [**here**](https://docs.aws.amazon.com/lambda/latest/dg/build-pipeline.html)**.** Many serverless frameworks exist out there that make managing a function easy. One might ask why I simply didn't convert the function to one of existing serverless frameworks such as [Serverless Framework](https://serverless.com/framework/). I was automating the deployment of a function for an AI engineer that had little to no knowledge about the cloud and simply wanted his function to run. I did not want to force him to learn a new framework, and took this as an opportunity to learn more practical knowledge. After two days of struggling, I was able setup a complete end-to-end pipeline that deploys the function under a REST API endpoint in API Gateway.

# Why CodePipeline? How about Jenkins?

Don't get me wrong. There are plenty of CI tools out there, and we can't go on without paying tribute to the king: JenkinsCI. Personally, I love Jenkins. The plugin marketplace is amazing, the community is ever growing larger. The sheer flexibility and power that comes with Jenkins is intoxicating. But, with great power, comes great responsibility. The responsibility of delivering a highly-available, salable CI service falls squarely on the DevOps engineer's shoulders. Unless you want to stay cloud-agnostic or have a valid reason to use Jenkins, then its best to use the managed services offered, as it is cheaper and provides high SLA with no work on your part. 

# The Architecture

![Lambda CI/CD Diagram](/static/lambda-cicd.jpg "Lambda CI/CD Architecture")

The architecture outlined above will be how we will set up the pipeline.

# Source Control

Source control providers supported by Code Pipeline are Code Commit and GitHub. Code Pipeline relies on Web hooks for GitHub, so it might not work properly with GitHub Enterprise. In that case, you will have to rely on Code Pipeline continuously polling to pull the latest change. I recommend using Code Commit when you can, as it is integrated with the AWS ecosystem.

# Build

AWS's answer to Jenkins was to make Code Build. Code Build is a managed service that takes a source and runs the commands outlined in the provided buildspec.yml inside a container. The container comes installed with mainstream tech stacks, such as Java, Nodejs, Go, Python, and is based on Ubuntu so installing build-time dependencies are easy. The buildspec.yml could be part of the project and checked in as part of the code, or could be created separately when you make the build project in Code Build. 

# Deploy

There are many deployment targets, but I would say the most common are Code Deploy, Cloudformation, and S3. For our CI/CD pipeline, we will be using Cloudformation to create and execute a changeset. By using a changeset, we can observe the changed resources, roll-back in cases of failure, and use deployment strategies like canary and blue/green releases.

# The Glue

AWS Code Pipeline is service is managed service that simply ties all of these services together. It takes care of storing and passing artifacts between stages. It monitors when a change has occurred in source control and the completion status of each of the stages. The service relies on AWS S3 to store artifacts between stages, giving us the capability to examine the artifacts between stages and when the pipeline fails. 

# Gotchas, Hiccups and Oops

Read the guide first without doing anything. **Twice**. Then follow the guide to the letter. Modifying the process beforehand just to save some time will only lead more wasted time.

## Oops

After setting up the pipeline, I encountered an error saying "failure to execute changeset". I pulled my hair out trying to figure out the problem. Then I realized I had skipped a step and did not give the role being used by Cloudformation enough permissions to execute the changeset. **If the pipeline fails at the Deploy stage, check to make sure the correct permissions are set.** 

## Hiccups

The guide encourages to use SAM to build your application. I myself was not familiar with SAM, so I had some trouble dealing with the template. After following the guide, I was trying to convert the function from Nodejs to Python, as my goal was to figure out how to automate the deployment of a Python function. If you are not using a framework for serverless functions like the [Serverless framework](https://serverless.com/framework/) or [Zappa](https://github.com/Miserlou/Zappa), then you will have to use templates and implement a handler. For the purpose of the guide, I would stick with templates, but for real-world usage, I would seriously consider the use of serverless frameworks. [Serverless Framework](https://serverless.com/framework/) especially has grown in popularity and supports an array of languages and providers, and its plugin system is very handy and useful.

## Gotchas

The function I was trying to automate was a disorganized Python function that uses NLTK to provide an AI-based service. The code needed to run the function and the code for development of the AI were in the same repository, so the repository was populated with training data sets and unnecessary packages. I had to move the necessary code to a separate folder so I could upload a zip file with only the needed code. In the SAM template, I refereed to an S3 URL for the Code Uri. However, I realized that the changeset generated could not detect any changes, so it would fail reporting that there were changes to execute. The template could not pick up on the fact that there was a new zip file available at the S3 URL. Another gotcha is that application dependencies also need to be zipped along with the function code. If you don't use layers, the dependencies can really add to your unzipped file size.

# Conclusion

After following the guide and applying the same procedure for a Python function instead of Nodejs, I learned a lot of practical knowledge in automating vanilla lambda functions with application dependencies. However, I do recommend using serverless frameworks whenever possible, as they had plenty of time to mature and make it easier to automate everything.
