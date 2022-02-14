#! /bin/bash

# - Installing goswagger:
# dir=$(mktemp -d)
# git clone https://github.com/go-swagger/go-swagger "$dir" 
# cd "$dir"
# go install ./cmd/swagger
# mv $GOPATH/bin/swagger $GOPATH/bin/goswagger

# - Generate configuration files (and server):
# mkdir goswagger
# goswagger generate server -A tutorAPI -f api/swagger.yaml -t goswagger