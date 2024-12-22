package main

import (
	"log"
	"net/http"

	"github.com/aitorfernandez/roshambo/gateway/client"
	"github.com/aitorfernandez/roshambo/gateway/handler"
	"github.com/aitorfernandez/roshambo/gateway/loader"
	"github.com/aitorfernandez/roshambo/gateway/middleware"
	"github.com/aitorfernandez/roshambo/gateway/resolver"
	"github.com/aitorfernandez/roshambo/gateway/schema"
	"github.com/aitorfernandez/roshambo/pkg/env"
	"github.com/gorilla/mux"
	"github.com/graph-gophers/graphql-go"
	"github.com/rs/cors"
)

func die(err error) {
	log.Fatalf("gateway %s", err.Error())
}

func main() {
	var (
		c   *client.Client
		err error
		sch string
	)
	if sch, err = schema.String(); err != nil {
		die(err)
	}
	if c, err = client.New(); err != nil {
		die(err)
	}

	// Create the GraphQL handler.
	gql := handler.NewGraphQL(
		graphql.MustParseSchema(sch, resolver.New(c)),
		loader.New(c),
	)
	m := mux.NewRouter().StrictSlash(true)
	m.Use(middleware.Addr)
	m.Handle("/graphql", gql)

	var h http.Handler
	if env.MustHget("app", "env") == "dev" {
		m.Handle("/", handler.NewGraphiQL())

		corsWrapper := cors.New(cors.Options{
			AllowedMethods: []string{"GET", "POST"},
			AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
		})
		h = corsWrapper.Handler(m)
	} else {
		h = m
	}

	// Configure the HTTP server.
	srv := &http.Server{
		Addr:           env.MustHget("gateway", "addr"),
		Handler:        h,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}
	if err = srv.ListenAndServe(); err != nil {
		die(err)
	}
}
