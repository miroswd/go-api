package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/miroswd/go-api/configs"
	"github.com/miroswd/go-api/handlers"
)

func main() {
	err := configs.Load()

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	var PORT = configs.GetServerPort()

	fmt.Printf("🔥 The server is running on port %v", PORT)
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), r)
}
