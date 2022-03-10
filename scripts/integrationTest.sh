#! /bin/bash

# while ! echo exit | nc localhost $1 | grep HTTP > /dev/null; do
#     sleep 5
# done

until $(curl --output /dev/null --silent --head http://localhost:$1); do
    sleep 5
done

echo 'Server up'

cat tests/setup.sql | make -s exec > /dev/null

go test -v ./tests/integration

cat tests/teardown.sql | make -s exec > /dev/null
