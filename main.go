package main

import (
	"log"
	"net/http"
	"os"

	"github.com/davidfunk13/overwatch-companion/database"
	router "github.com/davidfunk13/overwatch-companion/router"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

const defaultPort = "3001"

func main() {
	env := os.Getenv("APP_ENV")

	if env != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}

	database.InitConnection()
	Access to fetch at '' from origin '' has been blocked by CORS policy: Response to preflight request doesn't pass access control check: No 'Access-Control-Allow-Origin' header is present on the requested resource. If an opaque response serves your needs, set the request's mode to 'no-cors' to fetch the resource with CORS disabled.

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://overwatch-companion.netlify.app"},
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
