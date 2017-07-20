#!/usr/bin/env bash
# How to test
echo "Test Hello.php"
echo -e ">>> \033[1mhttp://localhost:8080/example?graphql=query{hello(message:\"World\")}\033[0m"
echo -e ">>> \033[1mhttp://localhost:8080/example?graphql=query{bye}\033[0m\n"

echo "Test Calc.php"
echo -e ">>> \033[1mhttp://localhost:8080/example?graphql=mutation{sum(x:1,y:2)}\033[0m\n"

echo "Test User.php"
echo -e ">>> \033[1mhttp://localhost:8080/example?graphql=query{user(id:\"1\"){name}}\033[0m"
echo -e ">>> \033[1mhttp://localhost:8080/example?graphql=query{user(id:\"id\"){name}}\033[0m"
echo -e ">>> \033[1mhttp://localhost:8080/example?graphql=query{user(id:\"编号\"){name}}\033[0m\n"

# Run PHP example server on port 8080.
php -S localhost:8080 server.php
