package main

import (
	_ "embed"
	"log"
	"net/http"

	"odunlamizo/book-collection/internal/database"
	"odunlamizo/book-collection/internal/graphql"

	"github.com/graphql-go/handler"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	if err := database.Initialize(); err != nil {
		log.Fatal(err)
	}
	h := handler.New(&handler.Config{
		Schema:     &graphql.Schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})
	http.Handle("/graphql", h)
	f := http.FileServer(http.Dir("web/static"))
	http.Handle("/", f)
	http.ListenAndServe(":8080", nil)
}
