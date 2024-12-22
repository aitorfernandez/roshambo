package resolver

import "github.com/aitorfernandez/roshambo/gateway/client"

// Resolver is the entry point for all GraphQL operations.
type Resolver struct {
	client *client.Client
}

// New creates a new Resolver.
func New(c *client.Client) *Resolver {
	return &Resolver{c}
}
