//go:build tools

// Package tools pins build-time tool dependencies (the gqlgen code generator) so `go mod tidy`
// keeps them resolvable for `go run github.com/99designs/gqlgen generate`.
package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/vektah/gqlparser/v2"
)
