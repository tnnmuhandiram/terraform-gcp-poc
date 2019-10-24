package main

import (
	"fmt"
	"log"
	"net/http"

	routes "./routes"
)

func main() {

	server := &http.Server{
		Addr:    fmt.Sprintf(":8000"),
		Handler: routes.Initilize(),
	}
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	}
}
