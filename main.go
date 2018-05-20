package main

import (
	"log"
	"net/http"
)

const assetsDir = "assets"

func main() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(assetsDir))))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
