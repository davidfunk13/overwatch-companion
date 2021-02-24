package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/davidfunk13/overwatch-companion/auth"
	"github.com/davidfunk13/overwatch-companion/graph"
	"github.com/davidfunk13/overwatch-companion/graph/generated"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

// NewRouter establishes all handlers for api routes
func NewRouter() *mux.Router {
	//get the current environment
	env := os.Getenv("APP_ENV")

	//create new router
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	origin := os.Getenv("ALLOWED_ORIGIN")

	//wrap all requests in cors handler
	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{origin},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	//auth middleware for all routes
	// r.Use(auth.JWTMiddleware.Handler)

	//graphql server
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	//look into each line of this
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				fmt.Println("AHHHHHHH")
				return r.Host == "https://overwatch-companion.netlify.app"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	//if you're in development
	if env != "production" {
		fmt.Println("You are in development env.", env)
		//set up the playground for graph queries
		r.Handle("/dev", playground.Handler("GraphQL playground", os.Getenv("GRAPH_SERVER")))
	}

	//serve graphql server at api endpoint.
	api.Use(auth.JWTMiddleware.Handler)
	api.Handle("/graph", srv)

	//we should do a profile endpoint with the auth0 management api here.

	//think about how we could use the scope funcits
	return r
}
