package resolver

import (
	"github.com/estebanborai/go-spacex-graphql/src/service"
)

// Resolver represents the GraphQL resolver
type Resolver struct {
	services service.ServiceProvider
}

// NewResolver creates an instance of the GraphQL
// resolver
func NewResolver(services service.ServiceProvider) *Resolver {
	return &Resolver{
		services: services,
	}
}
