package handler

import (
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

// GraphQL struct handles GraphQL API request over HTTP.
type GraphQL struct {
	handler *relay.Handler
}

// NewGraphQL creates a new GraphQL handler.
func NewGraphQL(s *graphql.Schema) *GraphQL {
	return &GraphQL{
		handler: &relay.Handler{
			Schema: s,
		},
	}
}

// ServeHTTP handler wraps ServeHTTP relay handler.
func (gql GraphQL) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	gql.handler.ServeHTTP(w, r)
}
