#!/usr/bin/env bash

resp=$(curl -q -s -X POST \
  -H "Content-Type: application/json" \
  -d '{"query": "query { hello(message: \"world\") }"}' \
  http://localhost:8080/graphql)
[[ "$resp" == "{\"data\":{\"hello\":\"Hello world!\"}}" ]] \
  && result="\033[32mOK\033[0m" || result="\033[31mFAIL\033[0m"
echo -e "$resp $result"

resp=$(curl -q -s -X POST \
  -H "Content-Type: application/json" \
  -d '{"query": "query { bye }"}' \
  http://localhost:8080/graphql)
[[ "$resp" == "{\"data\":{\"bye\":\"Goodbye!\"}}" ]] \
  && result="\033[32mOK\033[0m" || result="\033[31mFAIL\033[0m"
echo -e "$resp $result"

resp=$(curl -q -s -X POST \
  -H "Content-Type: application/json" \
  -d '{"query": "mutation { sum(x: 1, y: 2) }" }' \
  http://localhost:8080/graphql)
[[ "$resp" == "{\"data\":{\"sum\":3}}" ]] \
  && result="\033[32mOK\033[0m" || result="\033[31mFAIL\033[0m"
echo -e "$resp $result"

resp=$(curl -q -s -X POST \
  -H "Content-Type: application/json" \
  -d '{"query": "query{user(id:\"1\"){name}}" }' \
  http://localhost:8080/graphql)
[[ "$resp" == "{\"data\":{\"user\":{\"name\":\"+86-13888888888\"}}}" ]] \
  && result="\033[32mOK\033[0m" || result="\033[31mFAIL\033[0m"
echo -e "$resp $result"

resp=$(curl -q -s -X POST \
  -H "Content-Type: application/json" \
  -d '{"query": "query{user(id:\"id\"){name}}" }' \
  http://localhost:8080/graphql)
[[ "$resp" == "{\"data\":{\"user\":{\"name\":\"Name\"}}}" ]] \
  && result="\033[32mOK\033[0m" || result="\033[31mFAIL\033[0m"
echo -e "$resp $result"

resp=$(curl -q -s -X POST \
  -H "Content-Type: application/json" \
  -d '{"query": "query{user(id:\"编号\"){name}}" }' \
  http://localhost:8080/graphql)
[[ "$resp" == "{\"data\":{\"user\":{\"name\":\"名字\"}}}" ]] \
  && result="\033[32mOK\033[0m" || result="\033[31mFAIL\033[0m"
echo -e "$resp $result"
