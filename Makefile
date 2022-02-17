NAME := tutor

DOCKER-COMPOSE := COMPOSE_PROJECT_NAME=$(NAME) docker-compose
DOCKER-COMPOSE-PATH := ./config/docker-compose.yaml

SWAGGER-SPEC-PATH := ./docs/swagger.yaml
SWAGGER-MD-PATH := ./docs/swagger.md

all:		build up

build:
			$(DOCKER-COMPOSE) -f $(DOCKER-COMPOSE-PATH) build

up:
			$(DOCKER-COMPOSE) -f $(DOCKER-COMPOSE-PATH) up

down:
			$(DOCKER-COMPOSE) -f $(DOCKER-COMPOSE-PATH) down --remove-orphans

delete:
			$(DOCKER-COMPOSE) -f $(DOCKER-COMPOSE-PATH) down --volumes

re:			down all

install:
			go install ./cmd/$(NAME)-server/

run:		install
			./tutor-server --host 0.0.0.0 --port 5000

gen:		gen.ent gen.swag

gen.ent:
			go generate ./internal/ent

gen.swag:
			go generate ./internal/goswagger/restapi

validate:
			swagger validate $(SWAGGER-SPEC-PATH)

markdown:
			swagger generate markdown -f $(SWAGGER-SPEC-PATH) --output $(SWAGGER-MD-PATH)
