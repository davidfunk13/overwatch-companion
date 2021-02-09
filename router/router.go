package router

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/davidfunk13/overwatch-companion/auth"
	"github.com/davidfunk13/overwatch-companion/graph"
	"github.com/davidfunk13/overwatch-companion/graph/generated"
	"github.com/davidfunk13/overwatch-companion/helpers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// NewRouter establishes all handlers for api routes
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	addHandlers(r)
	return r
}

func addHandlers(r *mux.Router) {
	env := os.Getenv("APP_ENV")

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	// This route is always accessible
	r.Handle("/api/public", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		message := "Hello from a public endpoint! You don't need to be authenticated to see this."
		helpers.SendResponseJSON(message, w, http.StatusOK)
	}))

	if env != "production" {
		fmt.Println("You are in development env.", env)
		r.Handle("/dev", playground.Handler("GraphQL playground", "/api/query"))
		r.Handle("/api/query", srv)
	} else {
		r.Handle("/api/query", negroni.New(
			negroni.HandlerFunc(auth.JWTMiddleware.HandlerWithNext),
			negroni.Wrap(srv)))
	}

	// Mock private route
	r.Handle("/api/private", negroni.New(
		negroni.HandlerFunc(auth.JWTMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			message := "Hello from a private endpoint! You need to be authenticated to see this."
			helpers.SendResponseJSON(message, w, http.StatusOK)
		}))))

	// This route is only accessible if the user has a valid access_token with the read:messages scope
	// We are chaining the jwtmiddleware middleware into the negroni handler function which will check
	// for a valid token and scope.

	// Mock private router with permissions
	r.Handle("/api/private-scoped", negroni.New(
		negroni.HandlerFunc(auth.JWTMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeaderParts := strings.Split(r.Header.Get("Authorization"), " ")
			token := authHeaderParts[1]

			hasScope := auth.CheckScope("read:messages", token)

			if !hasScope {
				message := "Insufficient scope."
				helpers.SendResponseJSON(message, w, http.StatusForbidden)
				return
			}
			message := "Hello from a private endpoint! You need to be authenticated to see this."
			helpers.SendResponseJSON(message, w, http.StatusOK)
		}))))
}
