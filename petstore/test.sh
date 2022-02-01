#!/bin/bash

########################################################
# test.sh is a wrapper to execute all tests for
# petstore.
########################################################

echo "Petstore Loneliness 2000 Tests"
echo "Setup..."
soda drop -e test -c ./database.yml -p ./migrations
soda create -e test -c ./database.yml -p ./migrations
soda migrate -e test -c ./database.yml -p ./migrations
echo "Running Tests..."
go test ./... -v -count=1