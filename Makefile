NAME := tutor

all:		build up

build:
			docker-compose -f ./config/docker-compose.yaml build

up:
			docker-compose -f ./config/docker-compose.yaml up

down:
			docker-compose -f ./config/docker-compose.yaml down --remove-orphans

re:			down all

delete:
			docker-compose -f ./config/docker-compose.yaml down --volumes

run:
			go run ./cmd/$(NAME)-server/main.go --host 0.0.0.0 --port 5000

install:
			go install ./cmd/$(NAME)-server/

gen:		gen.ent gen.swag

gen.ent:
			go generate ./internal/ent

gen.swag:
			go generate ./internal/goswagger/restapi

validate:
			swagger validate ./docs/swagger.yaml

markdown:
			swagger generate markdown --output ./docs/swagger.md -f ./docs/swagger.yaml
