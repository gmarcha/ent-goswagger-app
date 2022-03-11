#! /bin/bash

while ! echo exit | nc localhost 5000 | grep HTTP > /dev/null; do
    sleep 5
done

# until $(curl --output /dev/null --silent --head http://localhost:$1); do
#     sleep 5
# done

echo 'Server up'

cat tests/setup.sql | docker-compose -p tutor \
    -f ./config/docker-compose.yaml -f ./config/docker-compose.ci.yaml --env-file ./config/.env.ci \
    exec -T -e POSTGRES_PASSWORD=tutor postgres psql -U tutor -d tutor > /dev/null

go test -v ./tests/integration

cat tests/teardown.sql | docker-compose -p tutor \
    -f ./config/docker-compose.yaml -f ./config/docker-compose.ci.yaml --env-file ./config/.env.ci \
    exec -T -e POSTGRES_PASSWORD=tutor postgres psql -U tutor -d tutor > /dev/null
