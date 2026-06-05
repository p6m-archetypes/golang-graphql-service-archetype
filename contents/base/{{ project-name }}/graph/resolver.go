package graph

//go:generate go run github.com/99designs/gqlgen generate

// Resolver is the root resolver. After running 'make generate',
// implement the generated QueryResolver and MutationResolver interfaces.
type Resolver struct{}
