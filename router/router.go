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
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

// CORSRouterDecorator applies CORS headers to a mux.Router
type CORSRouterDecorator struct {
	R *mux.Router
}

// ServeHTTP wraps the HTTP server enabling CORS headers.
// For more info about CORS, visit https://www.w3.org/TR/cors/
func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, YourOwnHeader")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}

	c.R.ServeHTTP(rw, req)
}

// NewRouter establishes all handlers for api routes
func NewRouter() *mux.Router {
	//get the current environment
	env := os.Getenv("APP_ENV")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
	})

	//create new router
	r := mux.NewRouter()

	//graphql server
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	//if you're in development
	if env != "production" {
		fmt.Println("You are in development env.", env)
		//set up the playground for graph queries
		r.Handle("/dev", playground.Handler("GraphQL playground", os.Getenv("GRAPH_SERVER")))
	}

	//if we're in production, instead put the graphql server in a Negroni instance with our jwt middleware.
	queryHandler := negroni.New()

	queryHandler.Use(negroni.HandlerFunc(auth.JWTMiddleware.HandlerWithNext))
	queryHandler.Use(negroni.Wrap(srv))
	queryHandler.Use(c)

	//serve graphql server at api endpoint.
	r.Handle(os.Getenv("GRAPH_SERVER"), queryHandler)

	//we should do a profile endpoint with the auth0 management api here.

	//think about how we could use the scope funcits
	return r
}
