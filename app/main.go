package main

import (
	"log"
	"net/http"
	"os"

	"github.com/CuriousHet/MiniFeed/graph"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {
	schemaData, err := os.ReadFile("graph/schema.graphql")
	if err != nil {
		log.Fatal("failed to read schema:", err)
	}

	schema := graphql.MustParseSchema(string(schemaData), &graph.Resolver{})

	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Println("Server running on http://localhost:8080/query")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
