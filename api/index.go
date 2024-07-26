package api

import (
	"net/http"

	"github.com/divyam234/instl/pkg/handlers"
	"github.com/gorilla/mux"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))
	router.HandleFunc("/{user}/{repo}/{os}", handlers.Installation)
	router.ServeHTTP(w, r)
}
