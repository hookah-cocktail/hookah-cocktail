package main

import (
	"log"
	"net/http"
	"os"
)

const assetsDir = "assets"

func main() {
	var port string
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s", port)
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(assetsDir))))
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
