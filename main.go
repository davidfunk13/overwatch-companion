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

// CORS Middleware
func CORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST")

			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token, Authorization")
			return
		}

		h.ServeHTTP(w, r)
	})
}
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
	http.Handle("/query", CORS(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
