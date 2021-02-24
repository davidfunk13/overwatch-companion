package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/davidfunk13/overwatch-companion/auth"
	"github.com/davidfunk13/overwatch-companion/graph"
	"github.com/davidfunk13/overwatch-companion/graph/generated"
	"github.com/davidfunk13/overwatch-companion/helpers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// NewRouter establishes all handlers for api routes
func NewRouter() *mux.Router {
	//get the current environment
	env := os.Getenv("APP_ENV")

	//create new router
	r := mux.NewRouter()

	origin := os.Getenv("ALLOWED_ORIGIN")

	fmt.Printf("What is this: \n %s", origin)

	//wrap all requests in cors handler
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{origin},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	//create a sub router for the graph api so we can protect only it with a jwt
	api := r.PathPrefix("/api").Subrouter()

	//graphql server handler
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	//if you're in development
	if env != "production" {
		fmt.Print(" ---------------------------- \n WELCOME TO DEVELOPMENT MODE \n ---------------------------- \n")

		//set up the playground for graph queries @ /dev without any auth.
		r.Handle("/dev", playground.Handler("GraphQL playground", os.Getenv("GRAPH_SERVER")))
	}

	//if you are in production, use a jwt on the api subrouter only
	if env == "production" {
		api.Use(auth.JWTMiddleware.Handler)
	}

	//serve the graph server at the api subrouter /graph.
	api.Handle("/graph", srv)

	//Heres a public route.
	r.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		message := "This is a public, restful api endpoint that will be used to return statstics from another api."
		helpers.SendResponseJSON(message, w, 200)
	})

	//we should do a profile endpoint with the auth0 management api here.

	//think about how we could use the scope funcits
	return r
}
