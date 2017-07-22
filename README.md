# graphql-examples

> Unified GraphQL examples written in Go, Node.js, PHP.

[![Build Status](https://travis-ci.org/WindomZ/graphql-examples.svg?branch=master)](https://travis-ci.org/WindomZ/graphql-examples)

## Purpose

- **Unified** GraphQL example template.
- **Support** multiple languages: Go, Node.js, PHP.
- **Easy** to learn primary GraphQL.
- **Test** compatibility.
- **HTTP** services.

## Template

Define a schema that is the basis for all queries: 
```
// add a User type
type User {
  id: String!
  name: String
}

// define a schema
type Query {
  hello(message: String): String
  bye: String
  user(id: String!): User
}
type Mutation {
  sum(x: Int, y: Int): Int
}
```

The example queries on the above schema would be:

Query `hello`: 
```
query Query {
  hello(message: "world")
}
```

Query `bye`: 
```
// test root values
query Query {
  bye
}
```

Query `user`: 
```
// test numbers and some symbols
query Query {
  user(id:"1") {
    name
  }
}
// test english strings
query Query {
  user(id:"id") {
    name
  }
}
// test chinese strings
query Query {
  user(id:"编号") {
    name
  }
}
```

Query `sum`: 
```
mutation Calc {
  sum(x: 1, y: 2)
}
```

## Related

- [graphql](https://github.com/facebook/graphql) Official documents
- [graphql-js](https://github.com/graphql/graphql-js) + [express-graphql](https://github.com/graphql/express-graphql)
- [graphql-go](https://github.com/graphql-go/graphql)
- [graphql-php](https://github.com/webonyx/graphql-php)

## Contributing

Welcome to pull requests, report bugs, suggest ideas and discuss **graphql-examples**, 
i would love to hear what you think about **graphql-examples** on [issues page](https://github.com/WindomZ/graphql-examples/issues).

If you like it then you can put a :star: on it.

## Todo

- [ ] Star Wars example template.
- [ ] HTTPS services.
- [ ] RPC services.

## License

[MIT](https://github.com/WindomZ/graphql-examples/blob/master/LICENSE)
