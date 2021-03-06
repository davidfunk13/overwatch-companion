package main

/*
to recompile all generated files with any changes to schemas, run
`go run github.com/99designs/gqlgen generate`
*/
import (
	"fmt"
	"log"
	"net/http"
	"os"

	router "github.com/davidfunk13/overwatch-companion/router"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const defaultPort = "3001"

var host string

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

	r := router.NewRouter()

	http.Handle("/", r)

	if env != "production" {
		host = "127.0.0.1"
		fmt.Printf(" +++ connect to http://localhost:%s/dev for GraphQL playground \n +++ Happy Hacking!\n", port)
	} else {
		host = ""
	}

	log.Fatal(http.ListenAndServe(host+":"+port, r))
}
