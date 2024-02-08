package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func BookRoutes() chi.Router {
	r := chi.NewRouter()

	//http.HandlerFunc functions can be fired from here

	//bookHandler.ListBooks ... etc are http.HandlerFunc functions

	// r.Get("/", bookHandler.ListBooks)
	// r.Post("/", bookHandler.CreateBook)
	// r.Get("/{id}", bookHandler.GetBooks)
	// r.Put("/{id}", bookHandler.UpdateBook)
	// r.Delete("/{id}", bookHandler.DeleteBook)

	return r
}

func main() {

	//creating a new router

	router := chi.NewRouter()

	//to use middlewares

	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RedirectSlashes)
	router.Use(middleware.Timeout(time.Second * 60))

	// request handling

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	// sub routing

	router.Mount("/books", BookRoutes())

	//listening server at 3000 port

	log.Fatal(http.ListenAndServe(":3000", router))

}
