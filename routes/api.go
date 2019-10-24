package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	gcp "github.com/tnnmuhandiram/terraform-gcp-poc/app/gcp"
)

//New http handler
func Initilize() http.Handler {
	routes := mux.NewRouter()

	routes.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Terraform POC Up and Running")
	})

	routes.HandleFunc("/gcp/compute/create", gcp.ComputeCreate).Methods("POST")
	// routes.HandleFunc("/users/{id}", user.Show).Methods("GET")
	// routes.HandleFunc("/users", user.AllUsers).Methods("GET")

	// route.Use(middleware)
	return routes
}
