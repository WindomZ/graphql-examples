#!/usr/bin/env bash
# How to test
echo -e "Test with GET: \033[1mhttp://localhost:8080/hello?query={hello}\033[0m"

echo -e "Test with GET: \033[1mhttp://localhost:8080/user?query={user(id:\"1\"){name}}\033[0m"
echo -e "Test with GET: \033[1mhttp://localhost:8080/user?query={user(id:\"id\"){name}}\033[0m"
echo -e "Test with GET: \033[1mhttp://localhost:8080/user?query={user(id:\"编号\"){name}}\033[0m"

# Run go example server on port 8080.
go run main.go