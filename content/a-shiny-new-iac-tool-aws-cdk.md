+++
title = "A Shiny New IaC Tool: AWS CDK"
date = 2020-04-17
+++

AWS has been continuously improving its cloud services day by day, and has never failed to amaze me. My website is currently hosted on AWS utilizing CodePipeline and static sites with S3, and I'm loving it. With the rise of IaC and DevOps, I barely have to touch the AWS console anymore to provision my infrastructure. AWS's answer to IaC was to build Cloudformation. I personally have barely used Cloudformation, as I felt a massive headache coming on every time I looked at the hundreds of lines of YAML code that must be written to achieve something significant. Of course, another very popular option is Terraform, which in my very biased opinion, is the de facto tool to provision infrastructure. Terraform supports multiple cloud providers like AWS, GCP, and Azure, and you can even support a whole new provider by writing a plugin yourself. The only problem with Terraform is the fact that you can't use a full blown programming language and must use HCL, Hashicorp's custom language. Naturally, I got very excited hearing about AWS CDK, where I could use the full blown power of a programming language like Python, Java, or Typescript to provision my infrastructure. Let's dive in!

# Initial Impressions

The very first thing I did was to follow the [CDK workshop](https://cdkworkshop.com/). The workshop, at the time of this writing, has support for Typescript, Python, Java, and dotNET. I followed the Typescript portion of the workshop. A small sidenote, the CDK is created in Typescript, then transpiled to other object oriented languages using [jsii](https://github.com/aws/jsii). Since Typescript is fairly easy to learn in a few hours, simple to set up, and is also the "first class" language for the CDK, I decided to follow the Typescript portion of the workshop.

According to the [workshop](https://cdkworkshop.com/):

> The AWS CDK is a new software development framework from AWS with the sole purpose of making it fun and easy to define cloud infrastructure in your favorite programming language and deploy it using AWS CloudFormation.

The AWS CDK is an abstract layer over Cloudformation. You define Cloudformation resources as code, and use the CDK CLI to synthesize and deploy Cloudformation stacks. In essence, you can create a CDK app, which can contain 1 or more stacks and each stack is a Cloudformation stack that can be deployed. In each stack, there are multiple constructs, and in each construct, can be another construct or a low-level Cloudformation resource.

A powerful feature of CDK is that you can nest constructs and create your own abstract high level constructs. For example, in the CDK workshop, you will create a hit counter for an existing lambda function and api. To view the hit counter in the browser without logging onto the AWS console, we imported a construct from npm called [cdk-dynamo-table-viewer](https://www.npmjs.com/package/cdk-dynamo-table-viewer), which exposed an endpoint for us to view the hit counters of the different paths that we query in our api. We didn't have to write up any new code for us create a new endpoint with a lambda function to view the hit counter. This, I believe, is the most powerful feature of the CDK: to be able to define constructs and package them. It blew my mind to see that infrastructure can now be **defined as npm packages**. Now you can distribute and manage infrastructure the same way you would manage libraries and packages in an organization or enterprise. This means that individuals and teams can define high level constructs that others can simply just **_use_**. To explore further, I recommend reading the [developer docs](https://docs.aws.amazon.com/cdk/latest/guide/home.html), following the [workshop](com), and taking a look at [awesome-cdk](https://github.com/eladb/awesome-cdk).

Another great benefit to having high level constructs is that the CDK can grant read or read/write permissions with just a method. Plenty of constructs in the CDK offer APIs for common use cases, like granting read and write permissions to a DynamoDB table for a Lambda function without actually mucking around with IAM policies.

# Drawbacks

While experimenting with the AWS CDK, I noticed some drawbacks that I would like to talk about. A glaring drawback is drift detection and auto-healing. If you have used Terraform before, then you know that Terraform can automatically detect drift, and restore your infrastructure back to its intended state. However, the AWS CDK is essentially a tool that synthesizes Cloudformation templates and deploys them. All of the drawbacks of Cloudformation will be apparent while using the CDK, and that includes the lack of automatic drift detection and auto-healing. At the time of this writing, there is no straightforward way to restore drifted infrastructure without recreating the entire stack. Many times, recreating the stack is not possible for resources like Cognito and DynamoDB. At times like these, Terraform plans out changes to fix the detected drifts. I encourage you to try out an experiment. After following the workshop, just before you clean up, go the console and delete one of the resources you created a like a DynamoDB table or a Lambda Function. Go back to the terminal, and type `cdk diff` or `cdk deploy`. You'll see that the cdk won't detect that you deleted some resources, granted that you didn't modify the code in between consecutive `diff`s or `deploy`s. To detect drift, you must utilize the console or use the AWS CLI, which only report the what resources have drifted from the stack. AWS has yet to implement an auto-healing feature for Cloudformation.

A feature that I really loved while using Terraform is [workspaces](https://www.terraform.io/docs/state/workspaces.html). According to Terraform [docs](https://www.terraform.io/docs/state/workspaces.html):

> The default workspace might correspond to the "master" or "trunk" branch, which describes the intended state of production infrastructure. When a feature branch is created to develop a change, the developer of that feature might create a corresponding workspace and deploy into it a temporary "copy" of the main infrastructure so that changes can be tested without affecting the production infrastructure.

This _really_ helped out while I was working with delicate resources like Cognito and DynamoDB. At my workplace, [Knowt](https://knowt.io/) was migrating thousands of users to a new user pool. The main reason was that our architecture tightly coupled Cognito to our data, which made it very difficult to extend our identity management system to include social providers. This transition created waves throughout our entire backend, requiring us to modify most of the backend to support the new identity management system. As the most experienced AWS expert at the company, I was tasked with some of the more complex jobs at the company like defining AppSync Schema and modifying the existing AppSync resolvers. Utilizing Terraform workspaces, I was able to create separate infrastructure for different environments and developers so that production wasn't disturbed until the release. There were many, _many_ times where dev infrastructure was torn down and rebuilt **within minutes** with filler data, such as hundreds of accounts for a dev Cognito user pool. Hooray for automation!!! Using Terraform workspaces made managing mission critical infrastructure across environments a trivial task. Workspaces became a feature I came to love, and I looked to see if there was a replacement for the AWS CDK. The short answer is that there is **no direct replacement**.

Workspaces is a feature built directly into Terraform. Something like that doesn't exist directly in the AWS CDK, but you can emulate that functionality by using the CDK context. A great example of this is [repo](https://github.com/aws-samples/aws-dynamodb-enterprise-application). While the CDK CLI can't natively create something like workspaces, you can create new apps for each given environment by modifying the context, then accessing the context during runtime.

# Well, should we use it?

With only my initial experimentation, I cannot give an accurate assessment. However, it is a very promising project that I think can offer a viable solution to Terraform. I think the only thing holding AWS CDK back is its lack of auto-healing and proper drift detection. This feature is a must for many people, especially when you have developers of differing experience working together. Mistakes will be made, but can be easily fixed using Terraform, many times without even re-provisioning the entire stack, whereas in Cloudformation, you have to jump through many hoops to fix modified infrastructure. Until Cloudformation and CDK can support auto-healing, I believe Terraform will continue to stand at the top of IaC tools for provisioning infrastructure. However, if you are in a situation where you will Cloudformation definitely check out CDK, it will significantly improve your dev experience.

**Footnotes:**

1. I say auto-healing because I don't know a better word for it
2. Here are some resources to jumpstart your exploration of AWS CDK
    1. <https://github.com/cdk-patterns/serverless> is a bunch of serverless patterns written using CDK
    2. <https://github.com/aws-samples/aws-cdk-examples> lots and lots of examples using Java, Typescript, Python, and the rest of the supported languages
    3. <https://github.com/aws-samples/aws-cdk-changelogs-demo> is an awesome sample showcasing the full power of CDK by provisioning a Serverless (Lambda and Farget) with ElasticSearch, Redis, SNS, etc.
