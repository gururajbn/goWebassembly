package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("listening on 8000")
	http.ListenAndServe(":8000", http.FileServer(http.Dir(".")))
}
