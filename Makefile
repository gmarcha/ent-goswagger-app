NAME := tutor

all:		gen run

run:
			go run ./cmd/$(NAME)-server/main.go --port 5000

install:
			go install ./cmd/$(NAME)-server/

gen:		gen.ent gen.swag

gen.ent:
			go generate ./ent

gen.swag:
			go generate ./goswagger/...

validate:
			swagger validate ./api/swagger.yaml
