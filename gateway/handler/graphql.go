package handler

import (
	"net/http"

	"github.com/aitorfernandez/roshambo/gateway/loader"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

// GraphQL struct handles GraphQL API request over HTTP.
type GraphQL struct {
	handler *relay.Handler
	loaders loader.Map
}

// NewGraphQL creates a new GraphQL handler.
func NewGraphQL(s *graphql.Schema, ll loader.Map) *GraphQL {
	return &GraphQL{
		handler: &relay.Handler{
			Schema: s,
		},
		loaders: ll,
	}
}

// ServeHTTP handler wraps ServeHTTP relay handler.
func (gql GraphQL) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := gql.loaders.Attach(r.Context())
	gql.handler.ServeHTTP(w, r.WithContext(ctx))
}
