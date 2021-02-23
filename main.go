package main

import (
	"log"
	"net/http"
	"os"

	"github.com/davidfunk13/overwatch-companion/database"
	router "github.com/davidfunk13/overwatch-companion/router"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

const defaultPort = "3001"

func main() {
	env := os.Getenv("APP_ENV")
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ALLOWED_ORIGIN")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

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

	r := router.NewRouter()

	http.Handle("/", r)

	if env != "production" {
		log.Printf("welcome to dev mode. connect to http://localhost:%s/dev for GraphQL playground", port)
	}

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(headersOk, originsOk, methodsOk)(r)))
}
