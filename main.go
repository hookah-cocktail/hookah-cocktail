package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/heroku/x/hmetrics/onload"
)

const assetsDir = "assets"

func main() {
	var port string
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(assetsDir))))
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
