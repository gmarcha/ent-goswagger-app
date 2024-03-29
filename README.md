# ent-goswagger-app

A web API written in Go programming language with entgo and goswagger.

[![CI](https://github.com/gmarcha/ent-goswagger-app/actions/workflows/ci.yaml/badge.svg)](https://github.com/gmarcha/ent-goswagger-app/actions/workflows/ci.yaml)
[![Doc](https://github.com/gmarcha/ent-goswagger-app/actions/workflows/doc.yaml/badge.svg)](https://github.com/gmarcha/ent-goswagger-app/actions/workflows/doc.yaml)
[![Tree](https://github.com/gmarcha/ent-goswagger-app/actions/workflows/tree.yaml/badge.svg)](https://github.com/gmarcha/ent-goswagger-app/actions/workflows/tree.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gmarcha/ent-goswagger-app)](https://goreportcard.com/report/github.com/gmarcha/ent-goswagger-app)
[![Go Pkg](https://pkg.go.dev/badge/github.com/gmarcha/ent-goswagger-app?status.svg)](https://pkg.go.dev/github.com/gmarcha/ent-goswagger-app)

![Image](assets/setup.svg)

## Requirements

- A Go [workspace](https://go.dev/doc/gopath_code), with Make[^1] and [Go Swagger](https://goswagger.io/install.html) already setup.
- Install [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/).
- Create `./config/.env.dev` file next to `./config/.env.sample` file with valid credentials (or `./config/.env.prod` for deploying purpose).

## Usage

Use Makefile rules to build and run application (see comments in `Makefile`).

These rules run against a specific workflow specified with `ENV` variable.\
Different environments are `dev`, `ci` and `prod`, `dev` is default value if none is provided.\
*For example, `make delete all ENV=ci` uses continuous integration (CI) workflow.*

- `make` triggers `make all`;
- `make all` uses `make build` and `make up` internally;
- `make re` uses `make down` and `make all` internally;
- `make build` builds docker-compose images;
- `make up` starts docker-compose containers;
- `make down` stops docker-compose containers;
- `make delete` stops docker-compose containers and erases docker-compose volumes;
- `make regen` uses `make gen` and `make reload` internally;
- `make reload` restarts API container to update source code.

Use them to perform code generation with entgo and goswagger.
 
- `make gen` uses `make gen.ent`, `make gen.doc` and `make gen.swag` internally;
- `make gen.serv` uses `make gen.doc` and `make gen.swag` internally;
- `make gen.ent` generates entgo data models from data schemas in `./internal/ent/schema/`;
- `make gen.doc` generates a swagger specification used to code generation and as a documentation in `./docs/`
    (an other swagger specification should be configure with routes and authentication model in `./config/`);
- `make gen.swag` generates goswagger server code from swagger specification in `./docs/` (do not use it
    directly, use `make gen.serv` instead);

Other rules are used by application Dockerfile (in `./config/`) and Github Actions (in `./.github/workflows/`).

## Repository

Repository architecture follows some conventions:

- `/cmd/` contains our applications, i.e. main packages responsible to launch them
    (it is a generated file in our case, but it won't be overwritten by further generation so it can be edited);

- `/config/` contains all our configuration files:
  - an environment file for each workflow: `.env.dev`, `.env.ci` and `.env.prod`;
  - a swagger specification (`spec.yaml`) used to configure our API route paths and
      our authentication model (data model definitions should be leave empty, they are added
      in an other swagger specification which serves as a model to code generation and as a
      documentation);
  - a `Dockerfile` defining an image to our API (based on golang:alpine);
  - a `docker-compose.yaml` for each workflow, responsible to launch needed application services
      with their defined environment;

- `/docs/` contains documentation about the project:
  - a generated swagger specification (`swagger.yaml`) used to code generation and
      documentation purposes (do not edit);
  - a markdown conversion of the specification (`swagger.md`), a file tree of the project and
      a workflow explaining the API processes;

- `/internal/` contains all other Go packages used in our applications:
  - `/modules/` contains packages defining our API logic, i.e. our data services,
      route handlers initialisation and definition of them, grouped in a logical way
      (group by users or events, in opposition to a functionnal partitioning);
  - `/clients/` contains packages related to data sources used by our services,
      i.e. loading environment, our SQL database access and possibly other SQL or non-SQL databases;
  - `/ent/` contains generated code from entgo (data schemas defined in `/ent/schema/`
      serve as data models to code generation, rules defined in `/internal/ent/entc` manage
      code generation itself);
  - `/goswagger/` contains swagger generated server code (various configuration and
      module initialisation must be written in `/goswagger/restapi/configure_tutor.go`);
  - `/utils/` contains various useful packages.

- `/tests/` contains a mock-server (to mock authorization API) and Go integration tests to our API services;

- `/vendor/` contains project dependencies, i.e. other Go modules needed by our API (dependencies are included
  in source code to avoid breaking changes in modules and to skip modules downloading in workflows);

- `/scripts/` contains scripts used by `Makefile` to run tests, to generate documentation or to setup a directory;
    
- `/.github/workflows/` contains our Github Actions' workflows, i.e. automated repository tasks triggered on specific events.

## Roadmap

Application development follows a design-first approach. It uses two code generation tools and it can be separated in **three steps**.

1. **Data schemas are written** as Go code in `/internal/ent/schema` and **a swagger specification is written** in `/config/spec.yaml` (but data definitions should be leave empty, they will be erase anyway). These configuration files serve as models to our code generation tools. Data schemas correspond to our application data models and the specification defines all route paths (our handlers), with an authentication policy associate to them. Authentication uses OAuth 2.0, an authorization delegation protocol, to manage access to our application. It implements different scopes internally allowing filtering access with roles.

2. **Code and assets generation is performed** with entgo, a bash script and goswagger. First it is necessary to generate Go data models from ent schemas (it also generates an openapi specification containing data definitions in `/internal/ent/` folder). Then a new swagger specification should be generated by merging the configuration specification and the data definitions specification. The new specification is a read-only file belonging to `/docs/` and serving as an API reference and documentation. Finally this specification serves as a model to perform server code generation with goswagger.

3. **Handlers and authentication mechanisms must be implemented** by developers. Generic handlers are implemented as modules, i.e. packages which gather all functionalities (handlers, services) related to an object in the same folder, rather than splitting them trough different packages. Login and callback handlers have to be configured to perform an OAuth handshake, to validate user access and to create and sign a JWT (JSON Web Token) based on user identity. A middleware is responsible to enforce user identity and to check roles on each request by verifying the JWT. It allows user access depending on request scopes (defined in OAuth security definition).

Continuous integration (with Github Actions) allows to run code validation, code generation and testing suites in an automatic workflow. Tasks are configured as jobs to be perform on some events (as a `push` or `pull-request`).

## Tools

- **Go**\
  Golang is a modern compiled programming language with strong typing, low-level capabilities and a builtin code generation tool.\
  *Links to [documentation](https://go.dev/doc/), [specification](https://go.dev/ref/spec) and [project layout](https://github.com/golang-standards/project-layout) examples.*

- **Testify**\
  Testify is a set of packages providing many tools to make test writing easier.\
  It provides assertion, mock objects and testing suites capabilities.\
  *Links to [documentation](https://pkg.go.dev/testing), to Go [`test` command](https://pkg.go.dev/cmd/go#hdr-Test_packages) and to [testify](https://pkg.go.dev/github.com/stretchr/testify#section-readme) package.*

---

- **Entgo**\
  Entgo is an entity framework, a graph oriented ORM, used to perform database access.\
  It abstracts SQL tables and queries as Go objects and methods.\
  Data models are directly written in Go code as data schemas.\
  Ent uses theses schemas to generate code.\
  *Link to [documentation](https://entgo.io/docs/getting-started/).*

---

- **GoSwagger**\
  A tool to generate server code based on an openapi specification version 2.0.\
  This specification is often called a swagger specification, because of the name of the firm which developed it.\
  The swagger firm also developed tools to generate a specification from source code or server/client/cli code from a specification.\
  Goswagger is a Go implementation of these tools.\
  *Links to [documentation](https://goswagger.io/) and how to use [external models](https://goswagger.io/use/models/schemas.html#external-types) for code generation.*
  
- **OpenAPI _(or Swagger)_ specification**\
  An openapi specification is a json (or yaml) file documenting authentication model, route paths and data models of an API.\
  A specification can serve as a model to generate server code (as with goswagger), or it can be generated from comments in source code (as with gin-swagger). It is also possible to generate it from Go data schemas (as with entoas or ogent, from entgo).\
  *Links to [documentation](https://swagger.io/docs/specification/2-0/basic-structure/) and [specification](https://swagger.io/specification/v2/).*
  
- **SwaggerEditor**\
  A tool to write and visualize an openapi specification online.\
  *Link to [Swagger Editor](https://editor.swagger.io/).*

---

- **OAuth 2.0**\
  An authorization delegation protocol. Authentication in application depends on this authorization.\
  *Links to [documentation](https://oauth.net/2/), Go [package](https://pkg.go.dev/golang.org/x/oauth2), how to use OAuth to implement [user authentication](https://oauth.net/articles/authentication/) and how to define [scopes](https://curity.io/resources/learn/scope-best-practices/) in the API.*

- **JSON Web Token**\
  JWT represents a way to safely transmit data between different parties in an application.\
  Tokens are generated by the api on login or refresh operations if the user is authorized.\
  Then they are sent to clients and clients must include the access token in every request.\
  They are signed to prevent unwanted modifications, so it allows to perform user authentication.\
  *Links to unofficial [documentation](https://jwt.io/) and Go [package](https://pkg.go.dev/github.com/golang-jwt/jwt).*

---

- **Make**\
  Make is a build automation tool used to build executables files. It can also be used to manage a project.\
  We use a file called Makefile which contains rules. These rules function as scripts.\
  *Links to [documentation](https://www.gnu.org/software/make/manual/make.html) and [wikipedia](https://en.wikipedia.org/wiki/Make_(software)).*

- **Docker** and **Compose**\
  Docker is an os-level virtualization technology based on containers.\
  They allow us to run our application in an isolated, platform independent environment.\
  Containers are run from images, themselves build from a Dockerfile.\
  Compose is a container orchestrator, used to manage a group of containers.\
  It is used to run simultaneously all our application services.\
  *Links to [Docker documentation](https://docs.docker.com/get-started/overview/) and to [Compose manual](https://docs.docker.com/compose/).*

- **PostGreSQL _(or any SQL database)_**\
  A SQL relational database, running in a container in development phase or with a cloud provider in production.\
  *Links to [documentation](https://www.postgresql.org/docs/14/index.html) and [DockerHub](https://hub.docker.com/_/postgres).*

- **PGAdmin _(or an other DB administration tool)_**\
  PGAdmin is a graphical database administration tool.\
  It allows to directly manipulate or monitore a database in a web browser.\
  *Links to [documentation](https://www.pgadmin.org/docs/pgadmin4/latest/) and how to [deploy](https://www.pgadmin.org/docs/pgadmin4/latest/container_deployment.html) it in a container*

- **Redis**\
  Redis is a lightweight key-value database.\
  It allows to easy manage caches or data stores in our application.\
  *Links to [documentation](https://redis.io/documentation) and Go [package](https://pkg.go.dev/github.com/go-redis/redis/v8).*

- **Traefik Proxy**\
  Traefik Proxy is an edge router which serves as a reverse proxy.\
  It intercepts and routes incoming request to our application.\
  Configuration rules are defined in a `docker-compose.yaml` file.\
  Traefik automatically discovers our infrastructure and manages tls certificates.\
  *Link to [documentation](https://doc.traefik.io/traefik/).*

---

- **GitHub Actions**\
  GH-Actions are CI (Continuous Integration) tools.\
  They allow us to run automated processes triggered by specific events on our repository.\
  Processes can be linters, validators or testers among other things.\
  Events can be `push`, `pull-request` or issue creation for example.\
  *Links to [documentation](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions) and a Go workflow [example](https://medium.com/swlh/setting-up-github-actions-for-go-project-ea84f4ed3a40).*

- **yq**\
  yq is a powerful yaml processor used to manipulate json or yaml files.\
  It allows to add, update, remove blobs on files or to merge, convert files.\
  It uses its own expression format as argument to perform modifications.\
  *Links to [documentation](https://mikefarah.gitbook.io/yq/) and how to [install](https://github.com/mikefarah/yq/#install) it.*

- **Useful links**\
  *How to use [markdown](https://docs.github.com/en/get-started/writing-on-github/getting-started-with-writing-and-formatting-on-github/basic-writing-and-formatting-syntax) on Github ? to [makeareadme](https://www.makeareadme.com/)? to draw an application [workflow](https://asciiflow.com/#/)? to [record](https://asciinema.org/) a terminal session (and to [convert](https://github.com/marionebl/svg-term-cli/) it to an animated svg file)?*

## Author

[@gmarcha](https://github.com/gmarcha)

## License

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

[^1]: `sudo apt install make` or an other package management tool for other distros.
