#! /bin/bash

# - Installing goswagger:
# dir=$(mktemp -d)
# git clone https://github.com/go-swagger/go-swagger "$dir" 
# cd "$dir"
# go install ./cmd/swagger
# mv $GOPATH/bin/swagger $GOPATH/bin/goswagger

# - Generate configuration files (and server):
# mkdir goswagger
# goswagger generate server -A tutorAPI -f api/swagger.yaml -P models.Principal -t goswagger

# After first generation, you could want to move `./goswagger/cmd/tutor-server/main.go` to `./cmd/tutor-server/main.go`,
# to allow modifications on main to be persistant. Please add --exclude-main in go generate rule in ./goswagger/restapi/configure_tutor.go.