###################################################################################

NAME := tutor

# Use rules defined in this file to interact with the application.

###################################################################################

-include config/.env

DOCKER-COMPOSE := COMPOSE_PROJECT_NAME=$(NAME) docker-compose
DOCKER-COMPOSE-PATH := ./config/docker-compose.yaml

SWAGGER-SPEC-PATH := ./config/spec.yaml
SWAGGER-DOC-PATH := ./docs/swagger.yaml
SWAGGER-MD-PATH := ./docs/swagger.md

TREE := ./docs/tree

######################################################################################################
#
#	Docker compose rules:
#
#	- `build` builds docker-compose images;
#	- `up` runs one container for each service defined in docker-compose.yaml;
#	- `down` stops containers and default application network;
#	- `delete` stops containers and removes their associated volumes,
#		i.e. persistent data storage (as postgres database);
#	- `reload` restarts API container to update code changes;
#	- `all` performs build and up rules;
#	- `re` restarts all containers using down and all;
#	- `regen` performs gen and reload rules.
#
#		This setup serves development purpose.
#
######################################################################################################

all:		build up

re:			down all

build:
			$(DOCKER-COMPOSE) -f $(DOCKER-COMPOSE-PATH) build --build-arg PORT=$(PORT)

up:
			$(DOCKER-COMPOSE) -f $(DOCKER-COMPOSE-PATH) up

down:
			$(DOCKER-COMPOSE) -f $(DOCKER-COMPOSE-PATH) down --remove-orphans

delete:
			$(DOCKER-COMPOSE) -f $(DOCKER-COMPOSE-PATH) down --volumes

regen:		gen reload

reload:
			$(DOCKER-COMPOSE) -f $(DOCKER-COMPOSE-PATH) restart goswagger

###################################################################################
#
#	API rules:
#
#	- `install` performs API installation in GOPATH (i.e. /go with default value in container);
#	- `run` launches installed server with host and port defined in environment.
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
#	- `gen.doc` generates a swagger specification which serves as a model
#		to code generation and as documentation, from an other
#		swagger specification used to configure routes and authentication;
#	- `gen.swag` generates server code from swagger specification;
#	- `gen` performs `gen.ent`, `gen.doc` and `gen.swag` rules;
#	- `gen.serv` performs `gen.doc` and `gen.swag` rules.
#
###################################################################################

gen:		gen.ent gen.doc gen.swag

gen.serv:	gen.doc gen.swag

gen.ent:
			go generate ./internal/ent

gen.doc:
			bash scripts/docgen.sh $(SWAGGER-SPEC-PATH) $(SWAGGER-DOC-PATH)

gen.swag:
			go generate ./internal/goswagger/restapi

###################################################################################
#
#	Continuous Integration (CI) rules:
#
#	- `validate` validates a swagger specification against 2.0 version;
#	- `markdown` generates a markdown description of a swagger specification;
#	- `tree` prints repository architecture with a tree representation.
#
###################################################################################

validate:
			swagger validate $(SWAGGER-DOC-PATH)

markdown:
			swagger generate markdown -f $(SWAGGER-DOC-PATH) --output $(SWAGGER-MD-PATH)

tree:
			echo '$> ent-goswagger-app git:(main) ✗ tree' > $(TREE); tree >> $(TREE)

###################################################################################
#
#	Project setup rules:
#
#	- `setup.gomod` creates a go mod file in current directory
#		(allowing module dependency tracking);
#	- `setup.ent` runs project initialisation for ent,
#		creating a schema directory and a generate.go file in ./internal/ent;
#	- `setup.swag` generates a goswagger sample project with a main,
#		then moves it to ./cmd/tutor-server/ and adds --exclude-main
#		to go generate rule in ./internal/goswagger/restapi/configure_tutor.go;
#	- `setup` performs `setup.gomod`, `setup.ent` and `setup.swag` rules.
#
###################################################################################

# setup:		setup.gomod setup.ent setup.swag

# setup.gomod:
# 			go mod init github.com/gmarcha/ent-goswagger-app

# setup.ent:
# 			bash scripts/setup/entgo.sh

# setup.swag:
# 			bash scripts/setup/goswagger.sh

###################################################################################
