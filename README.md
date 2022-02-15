# ent-goswagger-app

A back end API written in Go programming language, with authentication and authorization delegation.

## Software Stack

The API development follows a design first approach: a Swagger Specification is written, and from this specification, server code will be generated.\
The specification functions as a documentation for the API, containing all routes, response formats, and authentication models.\
Code generation tool is used to generate boilerplate server code, allowing us to focus on implementation of api logic.

- **Go**\
  It is a modern compiled programming language, with strong typing, low-level capabilities and code generation tools.\
  *Link to [documentation](https://go.dev/doc/), [specification](https://go.dev/ref/spec) and [project layout](https://github.com/golang-standards/project-layout) examples.*

- **GoEnt**\
  Ent is an entity framework, a graph oriented ORM, used to perform database access.\
  It abstracts SQL tables and queries as Go objects and methods.\
  *Link to [documentation](https://entgo.io/docs/getting-started/).*

- **PostGreSQL _(or any SQL database)_**\
  A SQL relationnal database, which run in a container in development phase and with a cloud provider in production.\
  *Links to [documentation](https://www.postgresql.org/docs/14/index.html) and [DockerHub](https://hub.docker.com/_/postgres).*
  
- **GoSwagger**\
  A very powerful code generation tool, to generate server code from a Swagger specification (OAPIv2).\
  *Link to [documentation](https://goswagger.io/).*
  
- **Swagger**\
  A Swagger Specification is a json/yaml file which documents routes, data schemas, and possibly authentication model about an API.\
  A specification can be written to serve as a model to generate server code (as with goswagger), or it can be generated from comments in source code (as with gin-swagger). It is also possible to generate it from go data schemas (as with elk or entoas, from entgo).\
  *Links to [documentation](https://swagger.io/docs/specification/2-0/basic-structure/) and [specification](https://swagger.io/specification/v2/).*
  
- **SwaggerEditor**\
  A tool to write and visualize Swagger Specification online.\
  *Link to [Swagger Editor](https://editor.swagger.io/).*
  
- **OAuth 2.0**\
  An authorization delegation protocol. In this application, authentication depends on this authorization.\
  *Links to [documentation](https://oauth.net/2/) and how to use it to implement [user authentication](https://oauth.net/articles/authentication/).*
