#! /bin/bash

# This script allows to use data models generated from entgo with goswagger server code generation tool.
# It avoids coexistence of two data models in our application.
# It enforces rules (and so validation) defined in entgo data schemas.

########################################################################################################

# Script arguments:
# - $1: swagger specification used as a configuration file,
#           for routes and authentication model;
# - $2: swagger specification which serves as a model
#           to goswagger and as a documentation,
#           i.e. created with merging first spec with an other
#           spec generated from entgo data schemas.

########################################################################################################

# How to install yq (yaml processor) ?
# - sudo apt-get install -y yq
# - or directly as a go pkg: GO111MODULE=on go get github.com/mikefarah/yq/v4

########################################################################################################

# Function arguments:
# - $1: definition name (for example User or Event);
# - $2: documentation swagger specification.

injectDefinitions() {

    # Add x-go-type field to data schemas in openapi specification generated from entgo generation tool.
    # It allows goswagger to skip data model generation and to use data models from entgo.
    VAR=".components.schemas.$1.x-go-type = { \"type\": \"$1\", \"import\": { \"package\": \"github.com/gmarcha/ent-goswagger-app/internal/ent\" } } | ."
    yq -i -o json "$VAR" internal/ent/openapi.json

    # Add data models to swagger specification used by goswagger and as project documentation.
    # It is created by merging our configurated swagger specification with one generated by entgo.
    VAR="select(fileIndex==0).definitions.$1 = select(fileIndex==1).components.schemas.$1 | select(fileIndex==0)"
    yq -iP ea "$VAR" $2 internal/ent/openapi.json
}

########################################################################################################

# Three steps:
# - create a new swagger specification which serves as a model and a documentation,
#       from an other swagger specification used as a configuration file,
#       for example, create ./docs/swagger.yaml from ./config/spec.yaml;
# - modify it to add our data models generated from entgo, available in ./ent/ as openapi.json;
# - add the x-go-type field in our models, to allow goswagger to skip model generation
#       and to use entgo data models;
# - replace '#/components/schemas/Type' with '#/definitions/Type' to make our spec 2.0 compliant.

echo "# Code generated by an automated script, DO NOT EDIT. Edit '$1' instead." > $2
yq $1 >> $2
for def in $(yq -P .components.schemas internal/ent/openapi.json | grep '^[A-Za-z]' | tr -d ':'); do
    injectDefinitions $def $2
done
sed -i 's/components\/schemas/definitions/g' $2

########################################################################################################
