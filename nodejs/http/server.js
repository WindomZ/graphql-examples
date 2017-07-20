/**
 * Created by WindomZ on 17-7-20.
 */
'use strict';

const express = require('express');
const graphqlHTTP = require('express-graphql');

const { schema: Schema, root: RootValue } = require('../graphql/schema');

let app = express();

app.use(
  '/example',
  graphqlHTTP({
    schema: Schema,
    rootValue: RootValue,
    graphiql: false,
    pretty: true,
  })
);

app.listen(8080);
console.log('Running a GraphQL API server at localhost:8080/example');
