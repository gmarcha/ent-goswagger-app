//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

// Link to an example of how to use entc package:
// https://github.com/ent/ent/tree/master/examples/entcpkg
// https://github.com/ent/ent/blob/master/entc/gen/template.go

func main() {
	err := entc.Generate("./schema", &gen.Config{
		Templates: []*gen.Template{
			gen.MustParse(
				gen.NewTemplate("validate").
					ParseFiles("template/validate.tmpl"),
			),
		},
	})
	if err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
