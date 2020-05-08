## Go GraphQL

Interact with a GraphQL server within our Go-based programs

Focused on the data-retrieval side of GraphQL

Well, consider working with systems that handle hundreds of thousands, if not millions of requests per day. Traditionally, we would hit an API that fronts our database and it would return a massive JSON response that contains a lot of redundant information that we might not necessarily need.

If we are working with applications at a massive scale, sending redundant data can be costly and choke our network bandwidth due to payload size.

GraphQL essentially allows us to cut down the noise and describe the data that we wish to retrieve from our APIs so that we are retrieving only what we require for our current task/view/whatever.

**One important thing to note is that GraphQL is not a query language like our traditional SQL. It is an abstraction that sits in-front of our APIs and is not tied to any specific database or storage engine.**

This is actually really cool. We can stand up a GraphQL server that interacts with existing services and then build around this new GraphQL server instead of having to worry about modifying existing REST APIs.

Simple GraphQL server in Go: https://github.com/graphql-go/graphql