/**
 * Created by WindomZ on 17-7-20.
 */
'use strict';

const _ = require('lodash');
const { GraphQLSchema, GraphQLObjectType } = require('graphql');

const { fields: HelloFields, root: HelloRoot } = require('./fields/hello');

const schema = new GraphQLSchema({
  query: new GraphQLObjectType({
    name: 'RootQueryType',
    fields: _.assignIn({}, HelloFields),
  }),
});

const root = _.assignIn({}, HelloRoot);

module.exports = { schema, root };
