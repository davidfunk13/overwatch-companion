package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/davidfunk13/overwatch-companion/database"
	"github.com/davidfunk13/overwatch-companion/graph"
	"github.com/davidfunk13/overwatch-companion/graph/generated"
	_ "github.com/go-sql-driver/mysql"
)

const defaultPort = "3001"

func main() {
	port := os.Getenv("PORT")

	if port == "" {

		port = defaultPort
	}

	/*
		Migration
		UP migrate -path database/migration -database "mysql://root@/overwatch_companion" -verbose up
		DOWN migrate -path database/migration -database "mysql://root@/overwatch_companion" -verbose down
	*/

	database.InitConnection()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
