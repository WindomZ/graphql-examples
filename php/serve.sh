#!/usr/bin/env bash
# How to test
echo -e "Test with GET: \033[1mhttp://localhost:8080/?query=query{hello(message:%22HelloWorld%22)}\033[0m"
echo -e "Test with GET: \033[1mhttp://localhost:8080/?query=mutation{sum(x:2,y:2)}\033[0m"
echo ""

# Run go example server on port 8080.
php -S localhost:8080 serve.php
