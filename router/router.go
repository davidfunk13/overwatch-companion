package router

import (
	"fmt"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/davidfunk13/overwatch-companion/auth"
	"github.com/davidfunk13/overwatch-companion/graph"
	"github.com/davidfunk13/overwatch-companion/graph/generated"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

// NewRouter establishes all handlers for api routes
func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	addHandlers(r)
	return r
}

func addHandlers(r *mux.Router) {
	env := os.Getenv("APP_ENV")

	//graphql server
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	//if you're in development
	if env != "production" {
		fmt.Println("You are in development env.", env)
		//set up the playground for graph queries
		r.Handle("/dev", playground.Handler("GraphQL playground", os.Getenv("GRAPH_API")))
	}

	//if we're in production, instead put the graphql server in a Negroni instance with our jwt middleware.
	queryHandler := negroni.New(negroni.HandlerFunc(auth.JWTMiddleware.HandlerWithNext), negroni.Wrap(srv))

	//serve graphql server at api endpoint.
	r.Handle(os.Getenv("GRAPH_API"), queryHandler)

	//we should do a profile endpoint with the auth0 management api here.

	//think about how we could use the scope functionality.
}
