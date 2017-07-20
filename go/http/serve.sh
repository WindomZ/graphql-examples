#!/usr/bin/env bash
# How to test
echo "Test hello.go"
echo -e ">>> \033[1mhttp://localhost:8080/example?graphql=query{hello(message:\"World\")}\033[0m\n"

echo "Test calc.go"
echo -e ">>> \033[1mhttp://localhost:8080/example?graphql=mutation{sum(x:1,y:2)}\033[0m\n"

echo "Test user.go"
echo -e ">>> \033[1mhttp://localhost:8080/example?graphql=query{user(id:\"1\"){name}}\033[0m"
echo -e ">>> \033[1mhttp://localhost:8080/example?graphql=query{user(id:\"id\"){name}}\033[0m"
echo -e ">>> \033[1mhttp://localhost:8080/example?graphql=query{user(id:\"编号\"){name}}\033[0m\n"

# Run go example server on port 8080.
go run server.go
