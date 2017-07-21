#!/usr/bin/env bash

curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"query": "query { hello(message: \"world\") }"}' \
  http://localhost:8080/graphql

curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"query": "query { bye }"}' \
  http://localhost:8080/graphql

curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"query": "mutation { sum(x: 1, y: 2) }" }' \
  http://localhost:8080/graphql
