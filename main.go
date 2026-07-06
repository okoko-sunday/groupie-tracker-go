package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	fs := http.FileServer(http.Dir("static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("server running on port http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
}
