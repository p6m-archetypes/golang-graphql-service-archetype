package graph

//go:generate go run github.com/99designs/gqlgen generate

import "{{ module_path }}/internal/repository"

// Resolver is the root resolver; it holds the service's dependencies.
type Resolver struct {
	Store *repository.Store
}
