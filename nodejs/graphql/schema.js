/**
 * Created by WindomZ on 17-7-20.
 */
'use strict';

const _ = require('lodash');
const { GraphQLSchema, GraphQLObjectType } = require('graphql');

const { fields: HelloFields, root: HelloRoot } = require('./fields/hello');
const { fields: CalcFields, root: CalcRoot } = require('./fields/calc');

const schema = new GraphQLSchema({
  query: new GraphQLObjectType({
    name: 'RootQueryType',
    fields: _.assignIn({}, HelloFields),
  }),
  mutation: new GraphQLObjectType({
    name: 'RootMutationType',
    fields: _.assignIn({}, CalcFields),
  }),
});

const root = _.assignIn({}, HelloRoot, CalcRoot);

module.exports = { schema, root };
