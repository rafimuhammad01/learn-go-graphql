package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	postgres "github.com/rafimuhammad01/learn-go-graphql/db/sqlc"
	"github.com/rafimuhammad01/learn-go-graphql/graph"
	"github.com/rafimuhammad01/learn-go-graphql/graph/generated"
	"log"
	"net/http"
	"os"
)

const (
	defaultPort = "8080"
)

func main() {
	// Load Env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect To Postgre
	db, err := postgres.NewConn(
		postgres.ConnString(os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD") , os.Getenv("DB_NAME")),
	)

	if err != nil {
		log.Fatal(err)
	}

	postgresDB := postgres.New(db)

	// Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{postgresDB}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
