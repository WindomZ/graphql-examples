/**
 * Created by WindomZ on 17-7-20.
 */
'use strict';

const { GraphQLString } = require('graphql');

// Construct a fields
const fields = {
  hello: {
    type: GraphQLString,
    resolve() {
      return 'Hello world!';
    },
  },
  bye: {
    type: GraphQLString,
  },
};

// The root provides a resolver function for each API endpoint
const root = {
  hello: () => {
    return 'Should be replaced by schema..hello.resolve()!';
  },
  bye: () => {
    return 'Goodbye!';
  },
};

module.exports = { fields, root };
