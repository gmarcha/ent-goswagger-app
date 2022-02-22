#! /bin/bash

# Script arguments:
# - $1: swagger specification used to goswagger code generation;
# - $2: swagger specification which serves as a documentation,
#           i.e. created with merging first spec with an other
#           spec generated from entgo data schemas.

# How to install yq (yaml processor) ?
# - sudo apt-get install -y yq

# Function arguments:
# - $1: definition name (for example User or Event);
# - $2: documentation swagger specification.
injectDefinitions() {
    VAR="select(fileIndex==0).definitions.$1 = select(fileIndex==1).components.schemas.$1 | select(fileIndex==0)"
    yq -iP ea "$VAR" $2 internal/ent/openapi.json
}

# Three steps:
# - create a new swagger specification which serves as documentation, from an other swagger specification,
#       for example, create ./docs/swagger.yaml from ./config/spec.yaml (used by goswagger to generate code);
# - modify it to add our data models generated from entgo, available in ./ent/ as openapi.json;
# - replace '#/components/schemas/Type' with '#/definitions/Type' to make our spec 2.0 compliant.
yq $1 > $2
for def in $(yq -P .components.schemas internal/ent/openapi.json | grep '^[A-Za-z]' | tr -d ':'); do
    injectDefinitions $def $2
done
sed -i 's/components\/schemas/definitions/g' $2