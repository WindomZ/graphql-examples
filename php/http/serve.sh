#!/usr/bin/env bash
# How to test
echo -e "Test with GET: \033[1mhttp://localhost:8080/?query=query{hello(message:%22World%22)}\033[0m"
echo -e "Test with GET: \033[1mhttp://localhost:8080/?query=query{bye}\033[0m"
echo -e "Test with GET: \033[1mhttp://localhost:8080/?query=mutation{sum(x:2,y:2)}\033[0m"
echo ""

# Run PHP example server on port 8080.
php -S localhost:8080 server.php
