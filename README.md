# ent-goswagger-app

A web API written in Go programming language with entgo and goswagger.

[![Golang Lint](https://github.com/gmarcha/ent-goswagger-app/actions/workflows/golangci-lint.yaml/badge.svg)](https://github.com/gmarcha/ent-goswagger-app/actions/workflows/golangci-lint.yaml)
[![Swagger Validation](https://github.com/gmarcha/ent-goswagger-app/actions/workflows/swaggerci-validate.yaml/badge.svg)](https://github.com/gmarcha/ent-goswagger-app/actions/workflows/swaggerci-validate.yaml)
[![Project Tree](https://github.com/gmarcha/ent-goswagger-app/actions/workflows/treeci.yaml/badge.svg)](https://github.com/gmarcha/ent-goswagger-app/actions/workflows/treeci.yaml)

## Prerequesites

- A Go [workspace](https://go.dev/doc/gopath_code) with Make *(`sudo apt install make`)* and [Go Swagger](https://goswagger.io/install.html) already setup.
- Install [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/).
- Create `./config/.env` file next to `./config/.env.sample` file (with valid credentials).

## Usage

Use Makefile rules to build and run application (see comments in `Makefile`).

- `make` uses `make build` and `make up` internally;
- `make build` builds docker-compose images;
- `make up` starts docker-compose containers;
- `make down` stops docker-compose containers;
- `make re` uses `make down` and `make` internally;
- `make delete` stops docker-compose containers and erases docker-compose volumes.

Use them to perform code generation with entgo and goswagger.
 
- `make gen` uses `make gen.ent` and `make gen.swag` internally;
- `make gen.ent` generates entgo model code from data schemas in `./internal/ent/schema/`;
- `make gen.swag` generates goswagger server code from swagger specification in `./docs/`.

Other rules are used by application Dockerfile (in `./config/`) and Github Actions (in `./.github/workflows/`).

## Development

Application development follows a design-first approach. It uses two code generation tools, and it can be separated in three steps.

1. Entgo generates Go types and methods to perform database access from data schemas directly written in Go code.\
   It enforces validation based on schema rules, which are applied to database-level (as foreign keys or SQL constraints).

2. Go Swagger generates server code from a swagger specification using version 2.0.\
   All security definitions, route paths and model definitions are written in the specification.\
   It can use external data models with swagger extension to avoid multiple data models in the application.\
   Specification serves as an API documentation for other developers.
   
3. API routes are interfaces with `Handle` method which need to be implemented by developer, as authentication mechanism.

## Stack

- **Go**\
  Golang is a modern compiled programming language, with strong typing, low-level capabilities and code generation tools.\
  *Links to [documentation](https://go.dev/doc/), [specification](https://go.dev/ref/spec) and [project layout](https://github.com/golang-standards/project-layout) examples.*

- **OAuth 2.0**\
  An authorization delegation protocol. In this application, authentication depends on this authorization.\
  *Links to [documentation](https://oauth.net/2/) and how to use it to implement [user authentication](https://oauth.net/articles/authentication/).*

---

- **GoSwagger**\
  A very powerful code generation tool, to generate server code from a Swagger specification (OAPIv2).\
  *Links to [documentation](https://goswagger.io/) and how to use [external models](https://goswagger.io/use/models/schemas.html#external-types) for code generation.*
  
- **Swagger**\
  A Swagger Specification is a json/yaml file which documents routes, data schemas and possibly authentication model about an API.\
  A specification can be written to serve as a model to generate server code (as with goswagger), or it can be generated from comments in source code (as with gin-swagger). It is also possible to generate it from Go data schemas (as with elk or entoas, from entgo).\
  *Links to [documentation](https://swagger.io/docs/specification/2-0/basic-structure/) and [specification](https://swagger.io/specification/v2/).*
  
- **SwaggerEditor**\
  A tool to write and visualize Swagger Specification online.\
  *Link to [Swagger Editor](https://editor.swagger.io/).*

---

- **GoEnt**\
  Ent is an entity framework, a graph oriented ORM, used to perform database access.\
  It abstracts SQL tables and queries as Go objects and methods.\
  *Link to [documentation](https://entgo.io/docs/getting-started/).*

- **PostGreSQL _(or any SQL database)_**\
  A SQL relationnal database, running in a container in development phase or with a cloud provider in production.\
  *Links to [documentation](https://www.postgresql.org/docs/14/index.html) and [DockerHub](https://hub.docker.com/_/postgres).*
  
 - **Docker** and **Compose**\
  Docker is an os-level virtualisation technology based on container.\
  They allow us to run our application in an isolated, platform independent environment.\
  Compose is a container orchestrator, used to manage a group of containers.\
  *Links to [Docker documentation](https://docs.docker.com/get-started/overview/) and to [Compose manual](https://docs.docker.com/compose/).*
  
---
  
- **GitHub Actions**\
  GH-Actions are CI (Continuous Integration) tools. They allow us to run automated processes triggered by specific events on our repository.\
  Processes can be linters, validators or testers among other things.\
  Events can be `push`, `pull-request` or issue creation for example.\
  *Links to [documentation](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions) and a Go workflow [example](https://medium.com/swlh/setting-up-github-actions-for-go-project-ea84f4ed3a40).*

- **Make**\
  Make is a build automation tool used to build executables files. Furthermore, it can be used to manage a project.\
  We use a file called Makefile, which contains rules. These rules function as scripts.\
  *Links to [documentation](https://www.gnu.org/software/make/manual/make.html) and [wikipedia](https://en.wikipedia.org/wiki/Make_(software)).*

- **Useful links**\
  How to [makeareadme](https://www.makeareadme.com/) ? Or how to [record](https://asciinema.org/) a terminal session.

## Author

- [@gmarcha](https://github.com/gmarcha)

## License

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
