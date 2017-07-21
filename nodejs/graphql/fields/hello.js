/**
 * Created by WindomZ on 17-7-20.
 */
'use strict';

const { GraphQLString } = require('graphql');

// Construct a fields
const fields = {
  hello: {
    type: GraphQLString,
    args: {
      message: {
        type: GraphQLString,
      },
    },
    resolve: (root, { message }) => 'Hello ' + message + '!',
  },
  bye: {
    type: GraphQLString,
  },
};

// The root provides a resolver function for each API endpoint
const root = {
  hello: () => {
    return 'Should not be displayed!';
  },
  bye: () => {
    return 'Goodbye!';
  },
};

module.exports = { fields, root };
