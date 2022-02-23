#! /bin/bash

# - Installing goswagger:
# git clone https://github.com/go-swagger/go-swagger
# go install ./go-swagger/cmd/swagger

# - Generate configuration files (and server):
mkdir -p internal/goswagger
swagger generate server -A tutor -f docs/swagger.yaml -P models.Principal -t internal/goswagger

# After first generation, you could want to move `./internal/goswagger/cmd/tutor-server/main.go` to `./cmd/tutor-server/main.go`,
# to allow modifications on main to be persistant and to add `--exclude-main` in go generate rule in `./internal/goswagger/restapi/configure_tutor.go`.

mkdir -p cmd/tutor-server
mv internal/goswagger/cmd/tutor-server/main.go cmd/tutor-server/main.go
sed -i 's/models.Principal/models.Principal --exclude-main/g' internal/goswagger/restapi/configure_tutor.go
