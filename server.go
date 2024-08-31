package main

import (
	"log"
	"net/http"
	"os"
	"time"


	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"kreid.com/graphl-go/graph"
	"github.com/joho/godotenv"

	"github.com/gorilla/websocket"

	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	err := godotenv.Load(); if err != nil {
		log.Fatal("Error loading .env file")
	}


	Database := graph.Connect()


	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers:
	&graph.Resolver{DB: Database}}))

	srv.AddTransport(transport.POST{})

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
		srv.Use(extension.Introspection{})


	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
