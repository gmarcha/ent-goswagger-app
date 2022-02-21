NAME := tutor

-include config/.env

DOCKER-COMPOSE := COMPOSE_PROJECT_NAME=$(NAME) docker-compose
DOCKER-COMPOSE-PATH := ./config/docker-compose.yaml

SWAGGER-SPEC-PATH := ./config/specification.yaml
SWAGGER-MD-PATH := ./docs/swagger.md

######################################################################################################
#
#	Docker compose rules:
#
#	- `build` builds docker-compose images, so postgres:alpine and goswagger (based on golang:alpine);
#	- `up` runs $(NAME)_postgres_1 and $(NAME)_goswagger_1 containers (default names);
#	- `down` stops both containers and their default associated network ($(NAME)_default);
#	- `delete` stops containers and removes docker-compose volumes,
#		i.e. data persistent storiage as postgres database;
#	- `all` performs build and up rules;
#	- `re` restarts both containers using down and all.
#
#		This setup serves development purpose.
#
######################################################################################################

all:		build up

build:
			$(DOCKER-COMPOSE) -f $(DOCKER-COMPOSE-PATH) build --build-arg PORT=$(PORT)

up:
			$(DOCKER-COMPOSE) -f $(DOCKER-COMPOSE-PATH) up

down:
			$(DOCKER-COMPOSE) -f $(DOCKER-COMPOSE-PATH) down --remove-orphans

delete:
			$(DOCKER-COMPOSE) -f $(DOCKER-COMPOSE-PATH) down --volumes

re:			down all

###################################################################################
#
#	API rules:
#
#	- `install` performs API installation in GoPath (i.e. /go with default value);
#	- `run` launches installed server with port 5000 on 0.0.0.0 network interface.
#
#		Note: 0.0.0.0 interface is required for server to be reachable
#			  from outside the container.
#
###################################################################################

install:
			go install ./cmd/$(NAME)-server/

run:		install
			$(NAME)-server --host $(HOST) --port $(PORT)

###################################################################################
#
#	Code generation rules:
#
#	- `gen.ent` generates entgo source code from ent data schemas;
#	- `gen.swag` generates server code from swagger specification;
#	- `gen` performs `gen.ent` and `gen.swag` rules.
#
###################################################################################

gen:		gen.ent gen.swag

gen.ent:
			go generate ./internal/ent

gen.swag:
			go generate ./internal/goswagger/restapi

###################################################################################
#
#	Swagger rules:
#
#	- `validate` validates a swagger specification agaisnt 2.0 version;
#	- `markdown` generates a markdown description of a swagger specification.
#
###################################################################################

validate:
			swagger validate $(SWAGGER-SPEC-PATH)

markdown:
			swagger generate markdown -f $(SWAGGER-SPEC-PATH) --output $(SWAGGER-MD-PATH)
