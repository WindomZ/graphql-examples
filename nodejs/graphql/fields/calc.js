/**
 * Created by WindomZ on 17-7-20.
 */
'use strict';

const { GraphQLInt } = require('graphql');

// Construct a fields
const fields = {
  sum: {
    type: GraphQLInt,
    args: {
      x: {
        type: GraphQLInt,
      },
      y: {
        type: GraphQLInt,
      },
    },
    resolve: (root, { x, y }) => x + y,
  },
};

// The root provides a resolver function for each API endpoint
const root = {};

module.exports = { fields, root };
