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

/*

1. Read the request path.
2. Split the path into parts.
3. Validate the number of parts.
4. Extract the ID portion.
5. Convert the ID from string to int.
6. Find the matching artist.
7. Return a 404 if the artist doesn't exist.
8. Render the artist page.


Notice how we added validation. Professional code always checks assumptions before using data.



can we learn this approach als:
Approach 3 (Modern Go)

This is what you'll use in professional Go applications.

Starting with Go 1.22, the standard library introduced pattern-based routing.

You can write:

http.HandleFunc("GET /artist/{id}", ArtistHandler)

Then inside the handler:

id := r.PathValue("id")
*/