package main

import (
	"log"
	"net/http"

	"github.com/aitorfernandez/roshambo/gateway/client"
	"github.com/aitorfernandez/roshambo/gateway/handler"
	"github.com/aitorfernandez/roshambo/gateway/middleware"
	"github.com/aitorfernandez/roshambo/gateway/resolver"
	"github.com/aitorfernandez/roshambo/gateway/schema"
	"github.com/aitorfernandez/roshambo/pkg/env"
	"github.com/gorilla/mux"
	"github.com/graph-gophers/graphql-go"
)

func die(err error) {
	log.Fatalf("gateway %s", err.Error())
}

func main() {
	var (
		clt *client.Client
		err error
		sch string
	)
	if sch, err = schema.String(); err != nil {
		die(err)
	}
	if clt, err = client.New(); err != nil {
		die(err)
	}

	// Create the GraphQL handler.
	gql := handler.NewGraphQL(
		graphql.MustParseSchema(sch, resolver.New(clt)),
	)
	m := mux.NewRouter().StrictSlash(true)
	m.Use(middleware.Addr)
	m.Handle("/graphql", gql)
	if env.MustHget("app", "env") == "dev" {
		m.Handle("/", handler.NewGraphiQL())
	}

	// Configure the HTTP server.
	srv := &http.Server{
		Addr:           env.MustHget("gateway", "addr"),
		Handler:        m,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}
	if err = srv.ListenAndServe(); err != nil {
		die(err)
	}
}
