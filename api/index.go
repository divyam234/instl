package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/installer/instl/internal/pkg/handlers"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()
	router.Use(mux.CORSMethodMiddleware(router))
	router.HandleFunc("/{user}/{repo}/{os}", handlers.Installation)
	router.ServeHTTP(w, r)
}
