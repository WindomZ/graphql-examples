/**
 * Created by WindomZ on 17-7-21.
 */
'use strict';

const test = require('ava');
const axios = require('axios');

test.serial('http server hello pass', async t => {
  await axios
    .get('http://localhost:8080/graphql?query=query{hello(message:"world")}')
    .then(response => {
      t.is(JSON.stringify(response.data), '{"data":{"hello":"Hello world!"}}');
      t.pass();
    })
    .catch(error => t.fail(error));

  await axios
    .get('http://localhost:8080/graphql?query=query{bye}')
    .then(response => {
      t.is(JSON.stringify(response.data), '{"data":{"bye":"Goodbye!"}}');
      t.pass();
    })
    .catch(error => t.fail(error));
});

test.serial('http server calc pass', async t => {
  await axios
    .post('http://localhost:8080/graphql', {
      query: 'mutation { sum(x: 1, y: 2) }',
    })
    .then(response => {
      t.is(JSON.stringify(response.data), '{"data":{"sum":3}}');
      t.pass();
    })
    .catch(error => t.fail(error));
});
