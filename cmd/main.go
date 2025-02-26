package main

import (
	"log"
	"net/http"
)

func main() {
	// create instance of server
	port := "3000"
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})

	log.Printf("\n\nServer is running at https://localhost:%s\n", port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
