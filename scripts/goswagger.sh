#! /bin/bash

# - Installing goswagger:
# dir=$(mktemp -d)
# git clone https://github.com/go-swagger/go-swagger "$dir" 
# cd "$dir"
# go install ./cmd/swagger

# - Generate configuration files (and server):
mkdir -p goswagger
swagger generate server -A tutor -f docs/swagger.yaml -P models.Principal -t goswagger

# After first generation, you could want to move `./goswagger/cmd/tutor-server/main.go` to `./cmd/tutor-server/main.go`,
# to allow modifications on main to be persistant. Please add `--exclude-main` in go generate rule in `./goswagger/restapi/configure_tutor.go`.