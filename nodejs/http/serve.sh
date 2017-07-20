#!/usr/bin/env bash
# How to test
echo -e "Test with GET: \033[1mhttp://localhost:8080/example?query={hello}\033[0m"
echo -e "Test with GET: \033[1mhttp://localhost:8080/example?query={bye}\033[0m"
echo ""

# Run Node.js example server on port 8080.
node server.js
