package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hegner123/modulacms.com/internal/models"
	"github.com/hegner123/modulacms.com/internal/request"
	"github.com/hegner123/modulacms.com/internal/templates"
)

func main() {
	// create instance of server
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}
	mux.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("public/styles"))))

	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("public/js"))))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := request.GetContent(r.URL.RawPath)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
		}
		m, err := models.BuildRoot(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
		}
		err = templates.BuildPage(m, w, r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
		}

	})

	log.Printf("\n\nServer is running at http://localhost:%s\n", "3000")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
