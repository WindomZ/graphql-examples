#!/usr/bin/env bash
# How to hello.js
echo -e ">>> \033[1mhttp://localhost:8080/graphql?query=query{hello(message:\"world\")}\033[0m"
echo -e ">>> \033[1mhttp://localhost:8080/graphql?query=query{bye}\033[0m\n"

# Run Node.js example server on port 8080.
node server.js
