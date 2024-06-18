package main

import (
	"fmt"
	"net/http"

	"graphql-go-crud/mutation"
	"graphql-go-crud/queries"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    queries.Query,
		Mutation: mutation.Mutation,
	},
)

func main() {

	corsHandler := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}

	http.Handle("/graphql", corsHandler(&handler.Handler{
		Schema: &schema,
	}))

	fmt.Println("Server is running on port 8082")
	http.ListenAndServe(":8082", nil)
	fmt.Println("Server is running on port 80810")

}
