#! /bin/bash

# - Installing goswagger:
# dir=$(mktemp -d)
# git clone https://github.com/go-swagger/go-swagger "$dir" 
# cd "$dir"
# go install ./cmd/swagger
# mv $GOPATH/bin/swagger $GOPATH/bin/goswagger

# - Generate configuration files (and server):
# goswagger generate server -A tutorsAPI -f api/swagger.yaml -t goswagger