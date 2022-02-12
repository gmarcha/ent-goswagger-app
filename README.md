# ent-goswagger-app

A back end api, with authentication and authorization delegation.

## Software Stack

- **Go**\
  It is a modern compiled programming language, with strong typing, low-level capabilities and code generation tools.\
  *Link to [documentation](https://go.dev/doc/), [specification](https://go.dev/ref/spec) and [project layout](https://github.com/golang-standards/project-layout).*

- **PostGreSQL**\
  A SQL relationnal database, which run in a container in development phase and with a cloud provider in production.\
  *Links to [documentation](https://www.postgresql.org/docs/14/index.html) and [DockerHub](https://hub.docker.com/_/postgres).*

- **GoEnt**\
  Ent is an entity framework, a graph oriented ORM, used to perform database access.\
  It abstracts SQL tables and queries as Go objects and methods.\
  *Link to [documentation](https://entgo.io/docs/getting-started/).*
  
- **GoSwagger**\
  A very powerful code generation tool, to generate boilerplate server code from a Swagger specification (OAPIv2).\
  *Link to [documentation](https://goswagger.io/).*
  
- **Swagger**\
  An Open Api Specification is a json file which documents routes, data schemas, and furthermore, authentication model about an API.\
  Spec is generated from comments in source code (as with gin-swagger), or it is written to generate server code (as with go-swagger). It could even be generated from data models (as with elk or entoas, from entgo).\
  *Links to [documentation](https://swagger.io/docs/specification/2-0/basic-structure/) and [specification](https://swagger.io/specification/v2/).*
  
- **SwaggerEditor**\
  A tool to write and generate Swagger Specification online.\
  *Link to [Swagger Editor](https://editor.swagger.io/).*
  
- **OAuth 2.0**\
  An authorization delegation protocol. In this application, authentication depends on this authorization.\
  *Links to [documentation](https://oauth.net/2/) and how to use it to implement user [authentication](https://oauth.net/articles/authentication/).*
