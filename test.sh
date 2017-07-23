#!/usr/bin/env bash
echo -e "======\033[46;30mTest CURL testings...\033[0m======"
./shell/http_server_test.sh
echo -e "======\033[46;30mTest Go testings...\033[0m======"
go test -v ./...
echo -e "======\033[46;30mTest Node.js testings...\033[0m======"
yarn run test
echo -e "======\033[46;30mTest PHP testings...\033[0m======"
composer test
