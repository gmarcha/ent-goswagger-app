//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entoas"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

// Link to an example of how to use entc package:
// https://github.com/ent/ent/tree/master/examples/entcpkg
// https://github.com/ent/ent/blob/master/entc/gen/template.go

func main() {
	oas, err := entoas.NewExtension(
		entoas.DefaultPolicy(entoas.PolicyExclude),
	)
	if err != nil {
		log.Fatalf("creating entoas extension: %v", err)
	}
	err = entc.Generate("./schema", &gen.Config{
		Templates: []*gen.Template{
			gen.MustParse(
				gen.NewTemplate("validate").
					ParseFiles("template/validate.tmpl"),
			),
		},
	}, entc.Extensions(oas))
	if err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
