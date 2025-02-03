// cmd/main.go
package main

import (
	"log"
	"net/http"
	"todo-backend-golang/pkg/router"
)

func main() {
	r := router.SetupRouter()

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error starting server : %s", err)
	}
}
