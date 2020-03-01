package config

import (
	"github.com/estebanborai/go-spacex-graphql/src/resolver"
	"github.com/estebanborai/go-spacex-graphql/src/schema"
	"github.com/estebanborai/go-spacex-graphql/src/service"

	graphql "github.com/graph-gophers/graphql-go"
)

// NewGraphQLSchema generates a new resolver given a database
// connection and parses the schema provided by the "schema"
// package.
// Finally returns a pointer to the GraphQL schema
func NewGraphQLSchema(services service.ServiceProvider) *graphql.Schema {
	return graphql.MustParseSchema(schema.GetRootSchema(), resolver.NewResolver(services))
}
