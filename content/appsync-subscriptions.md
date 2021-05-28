+++
title = "The Complete Guide to AppSync Subscriptions"
date = 2020-06-15
[taxonomies]
tags = ["AWS", "Tutorial", "AppSync", "GraphQl"]
+++

Appsync Subscriptions are important for anyone that wants to develop a realtime app. However, I had a lot of trouble trying to understand and use subscriptions. The documentation provided plenty of help and laid down the foundations, but there were some concepts that were discovered through Googling, trial and error, and Stack Overflow. My hope in writing the following post is that you won't struggle as much as I did, and will avoid the mistakes I made while working with Appsync subscriptions.

**Before you go on, this post assumes that you have working knowledge of Graphql and (partially) of Appsync.**

## **TL;DR**

- You can control what gets sent from subscription by using subscription filters (defining parameters in subscriptions)
- Resolvers attached to subscriptions run only once when the client connects
- Even though subscription resolvers do not return any data to the client, they must resolve to returning type. This means that any non nullable fields in the returning type must be resolved.

# Why Subscriptions?

The first question to ask is if you really need subscriptions. If you want to receive updates, and your app or service is tolerant of missed updates, I would argue that a simple polling mechanism is better. A polling mechanism can scale better, is cheaper to use and can be customized to fit your exact needs. However, if you are sure that subscriptions is necessary, then you should be aware of a few things.

# Simple Subscriptions

Subscriptions in Appsync are deceptively simple. All that is really needed to create a subscriptions is to create the subscription field, provide a simple directive, and viola🎻, you have a subscription! Let's take a look at an example schema from the official [docs](https://docs.aws.amazon.com/appsync/latest/devguide/real-time-data.html):

```graphql
# ... rest of schema ommited for brevity

type Mutation {
	addPost(id: ID! author: String! title: String content: String url: String): Post!
	updatePost(id: ID! author: String! title: String content: String url: String ups: Int! downs: Int! expectedVersion: Int!): Post!
	deletePost(id: ID!): Post!
}

type Subscription {
	addedPost: Post # Add a subscription field
	@aws_subscribe(mutations: ["addPost"]) # add the graphql directive

	updatedPost: Post
	@aws_subscribe(mutations: ["updatePost"])

	deletedPost: Post
	@aws_subscribe(mutations: ["deletePost"])
}
```

The subscription `addedPost` will run **after** the mutation addPost has been run. The output of the mutation will be the input of the subscription, so **the type returned by the subscription must match the type returned by subscribed mutation**. In this case, both `addedPost` and `addPost` return a `Post` type. This is all that is needed to set up a simple subscription. However, we usually also want to control what updates we receive from subscriptions, as well as set up some security measures so that only specific clients can use subscriptions.

# Subscription Filters

We will handle the first case first, which is controlling what gets sent back from the subscriptions. Let's take a look at this schema that was presented in the [docs](https://docs.aws.amazon.com/appsync/latest/devguide/real-time-data.html):

```graphql
# ... rest of schema ommited for brevity

type Comment {
	# The id of the comment's parent event.
	eventId: ID!
	# A unique identifier for the comment.
	commentId: String!
	# The comment's content.
	content: String
	# Location where the comment was made
	location: String
}

type Event {
	id: ID!
	name: String
	where: String
	when: String
	description: String
}

type Mutation {
	commentOnEvent(eventId: ID!, location: String, content: String): Comment
}

type Subscription {
	# The subscription takes an argument to filter through events
	subscribeToEventComments(eventId: String!, location: String, content: String): Comment
	@aws_subscribe(mutations: ["commentOnEvent"])
}
```

The way we can control what gets sent back from subscriptions is by using **subscription filters**. I mentioned previously that the output of a mutation gets sent to a subscription, which then forwards the details to the client. We can "control" what gets sent back by making sure certain conditions are met. For example, using the schema above, `subscribeToEventComments` will send a `Comment` type to the client only when the `eventId` in the returned object is the same as the id that the client provides as a parameter. For example, if the client wants to subscribe to an event's comments with the `eventId` of `event1`, then the client will only receive that `Comments` pertaining to `event1` and not `event2`. I previously mentioned that the output of the subscribed mutation gets sent directly to the client, but with "filters", or provided parameters, Appsync will first check to make sure that value of the parameters matches what's in the field. This means that **all subscription parameters must exist as a field in the returning type**. This means that the type of the parameter must also match the type of the field in the returning object. For example, `eventId` must be of type `String`. It cannot be any other type. 

When creating subscriptions with parameters, you can also specify if you want the argument to be optional. In that case, the behavior is fairly straightforward: if an optional field is provided, then Appsync will use it to determine whether to send data back to the client. The only tricky behavior is if an optional argument is set as null; **Appsync will only send back data where the corresponding field is also null**. For example, in the schema above, if you provide an argument similar to `location: null`, then Appsync will only send back Comments when **the location field is null**.

# Subscription Resolvers

Besides subscriptions filters, we can also attach resolvers to control who can subscribe to the fields. For example, if we take a look at this schema from the official [docs](https://docs.aws.amazon.com/appsync/latest/devguide/security-authorization-use-cases.html#real-time-data):

```graphql
input CreateUserPermissionsInput {
    user: String!
    isAuthorizedForSubscriptions: Boolean
}

type Message {
    id: ID
    toUser: String
    fromUser: String
    content: String
}

type MessageConnection {
    items: [Message]
    nextToken: String
}

type Mutation {
    sendMessage(toUser: String!, content: String!): Message
    createUserPermissions(input: CreateUserPermissionsInput!): UserPermissions
    updateUserPermissions(input: UpdateUserPermissionInput!): UserPermissions
}

type Query {
    getMyMessages(first: Int, after: String): MessageConnection
    getUserPermissions(user: String!): UserPermissions
}

type Subscription {
    newMessage(toUser: String!): Message
        @aws_subscribe(mutations: ["sendMessage"])
}

input UpdateUserPermissionInput {
    user: String!
    isAuthorizedForSubscriptions: Boolean
}

type UserPermissions {
    user: String
    isAuthorizedForSubscriptions: Boolean
}

schema {
    query: Query
    mutation: Mutation
    subscription: Subscription
}
```

we want to make sure that users can only subscribe to conversations that pertain to them. We don't want random users to receive messages that are supposed to go to another user. We can attach a  simple resolver that checks the user's identity, and that the provided `toUser` value really is the user. The docs go into slightly more complex use cases and are more verbose, but do a pretty good job explaining how subscription revolvers can be used, so you can check that out [here](https://docs.aws.amazon.com/appsync/latest/devguide/security-authorization-use-cases.html#real-time-data).

However, what is not mentioned is that resolvers run **only once**. They run only once when the user calls the subscription. **There is no way to control what gets sent through subscriptions except through filters *during* the actual connection**. To clarify, we can use **resolvers to reject or accept a subscription connection to client**, however, **we cannot stop the subscription ourselves**, nor can we control what gets sent back to the client on case by case basis.

Another "gotcha" I encountered is that when you attach a resolver, you must resolve the type being returned! Even though the resolver attached to subscription does not actually send back any data to the client, we must still adhere the schema and resolve any mandatory fields of the returning type. For example, let's assume that the `Message` type in the schema above required that the `id` field be of type `String!`. This means that when we attach a resolver to our subscription, we must resolve to a `Message` type that has a valid value for the `id`. Keep in mind that **the subscription resolver will not send data back to the client**, and as such can resolve to a "fake" `Message`. We can provide any arbitrary value for the `id`, and it wouldn't matter.

# Conclusion

Appsync is a great service to bootstrap your service or app, but it does have its limitations. Subscriptions have really quirky behavior, and Appsync only provides limited control over subscriptions. For those looking for complete control of their schema, I would recommend to run their own Graphql service on top of scalable infrastructure like Lambda, Fargate or Kubernetes. Getting started with Appsync is very fast, but as your use cases evolve and become more complex, it might actually slow down developer velocity.

Hopefully you learned something you didn't know or resolved some of your doubts about Appsync in this article. Happy developing! 🎉