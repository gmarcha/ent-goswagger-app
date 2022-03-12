package ent

import (
	// Used to add module to dependencies (used by entc).
	_ "entgo.io/contrib/entoas"
)

//go:generate go run -mod=mod entc/entc.go
