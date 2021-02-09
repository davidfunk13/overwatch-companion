package main

import (
	"log"
	"net/http"
	"os"

	"github.com/davidfunk13/overwatch-companion/database"
	router "github.com/davidfunk13/overwatch-companion/router"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

const defaultPort = "3001"

func main() {
	env := os.Getenv("APP_ENV")

	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}

	database.InitConnection()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization"},
	})

	r := router.NewRouter()

	handler := c.Handler(r)

	http.Handle("/", r)

	if env != "production" {
		log.Printf("welcome to dev mode. connect to http://localhost:%s/dev for GraphQL playground", port)
	}

	log.Fatal(http.ListenAndServe(":"+port, handler))
}
