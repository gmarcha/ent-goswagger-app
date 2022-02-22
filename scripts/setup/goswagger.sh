#! /bin/bash

# - Installing goswagger:
# git clone https://github.com/go-swagger/go-swagger
# go install ./go-swagger/cmd/swagger

# - Generate configuration files (and server):
mkdir -p internal/goswagger
swagger generate server -A tutor -f config/spec.yaml -P models.Principal -t internal/goswagger

# After first generation, you could want to move `./goswagger/cmd/tutor-server/main.go` to `./cmd/tutor-server/main.go`,
# to allow modifications on main to be persistant. Please add `--exclude-main` in go generate rule in `./goswagger/restapi/configure_tutor.go`.