/**
 * Created by WindomZ on 17-7-20.
 */
'use strict';

const express = require('express');
const graphqlHTTP = require('express-graphql');
const { GraphQLSchema, GraphQLObjectType, GraphQLString } = require('graphql');

// Construct a schema
let schema = new GraphQLSchema({
  query: new GraphQLObjectType({
    name: 'RootQueryType',
    fields: {
      hello: {
        type: GraphQLString,
        resolve() {
          return 'Hello world!';
        },
      },
      bye: {
        type: GraphQLString,
      },
    },
  }),
});

// The root provides a resolver function for each API endpoint
let root = {
  hello: () => {
    return 'Should be replaced by schema..hello.resolve()!';
  },
  bye: () => {
    return 'Goodbye!';
  },
};

let app = express();

app.use(
  '/example',
  graphqlHTTP({
    schema: schema,
    rootValue: root,
    graphiql: false,
    pretty: true,
  })
);

app.listen(8080);
console.log('Running a GraphQL API server at localhost:8080/example');
