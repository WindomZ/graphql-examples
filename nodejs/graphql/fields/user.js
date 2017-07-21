/**
 * Created by WindomZ on 17-7-21.
 */
'use strict';

const { GraphQLObjectType, GraphQLString } = require('graphql');

class User {
  constructor(id, name) {
    this.id = id;
    this.name = name;
  }
}

let data = {
  '1': new User('1', '+86-13888888888'),
  id: new User('id', 'Name'),
  编号: new User('编号', '名字'),
};

const UserType = new GraphQLObjectType({
  name: 'user',
  isTypeOf: obj => Promise.resolve(obj instanceof User),
  fields: {
    id: { type: GraphQLString },
    name: { type: GraphQLString },
  },
});

// Construct a fields
const fields = {
  user: {
    type: UserType,
    args: {
      id: {
        type: GraphQLString,
      },
      name: {
        type: GraphQLString,
      },
    },
    resolve: (root, { id, name }) => data[id],
  },
};

// The root provides a resolver function for each API endpoint
const root = {};

module.exports = { fields, root };
