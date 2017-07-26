# graphql-examples

> Unified GraphQL examples written in Go, Node.js, PHP.

[![Build Status](https://travis-ci.org/WindomZ/graphql-examples.svg?branch=master)](https://travis-ci.org/WindomZ/graphql-examples)
[![Go Report Card](https://goreportcard.com/badge/github.com/WindomZ/graphql-examples)](https://goreportcard.com/report/github.com/WindomZ/graphql-examples)
[![Greenkeeper badge](https://badges.greenkeeper.io/WindomZ/graphql-examples.svg)](https://greenkeeper.io/)
[![Dependency](https://david-dm.org/WindomZ/url-generator.svg)](https://david-dm.org/WindomZ/url-generator)
[![styled with prettier](https://img.shields.io/badge/js_styled_with-prettier-brightgreen.svg)](https://github.com/prettier/prettier)
[![License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

## Purpose

- **Unified** GraphQL example template.
- **Multiple** languages: Go, Node.js, PHP.
- **Easy** to learn primary GraphQL.
- **Test** compatibility.
- **HTTP** services.

## Install

### Go
```bash
go get github.com/WindomZ/graphql-examples/...
```

### Node.js
```bash
npm install
# or
yarn
```

### PHP
```bash
composer install
```

## Usage

### Single

Each under the `go`, `nodejs`, `php` directory paths.

### Integrated

If you have `Go(go get)`, `PHP(composer)`, `Node.js(npm/yarn)` installed.

You can run and select service:
```
./run.sh
```
Run the test with all above languages:
```
./test.sh
```

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
// test English
query Query {
  user(id:"id") {
    name
  }
}
// test Chinese
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

Welcome to pull requests, report bugs, suggest ideas and discuss 
**graphql-examples** on [issues page](https://github.com/WindomZ/graphql-examples/issues).

If you like it then you can put a :star: on it.

## Roadmap

- [x] Support Golang.
- [x] Support Node.js.
- [x] Support PHP.
- [ ] Support Java.
- [x] Hello & Bye example template.
- [x] Calc(sum) example template.
- [x] User example template.
- [ ] Star Wars example template.
- [x] HTTP services.
- [ ] HTTPS services.
- [ ] RPC services.

## License

[MIT](https://github.com/WindomZ/graphql-examples/blob/master/LICENSE)
