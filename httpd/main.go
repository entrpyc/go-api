package main

import (
	"go-api/platform/newsfeed"
	"github.com/go-chi/chi/v5"
	"net/http"
	"encoding/json"
	"fmt"
	"go-api/httpd/handler"
)

func main() {
	port := ":3000"
	r := chi.NewRouter()
	feed := newsfeed.New()

	r.Get("/newsfeed", handler.NewsfeedGet(feed))

	r.Post("/newsfeed", handler.NewsfeedPost(feed))

	fmt.Println("Serving on " + port)
	http.ListenAndServe(port, r)
}