package routes

import (
	"fmt"
	"net/http"

	gcp "../app"
	"github.com/gorilla/mux"
)

//New http handler
func Initilize() http.Handler {
	routes := mux.NewRouter()

	routes.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Terraform POC Up and Running")
	})

	routes.HandleFunc("/gcp/compute/create", gcp.computeCreate).Methods("POST")
	// routes.HandleFunc("/users/{id}", user.Show).Methods("GET")
	// routes.HandleFunc("/users", user.AllUsers).Methods("GET")

	// route.Use(middleware)
	return routes
}
