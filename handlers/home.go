package handlers

import (
	"groupie/api"
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	artists, err := api.FetchArtists()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		log.Println("internal server error")
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	err = tmpl.Execute(w, artists)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		log.Printf("eror executing template: %v", err)
		return
	}

}





/*
func ArtistHandler(w http.ResponseWriter, r *http.Request){
	path := r.URL.Path

	part := strings.Split(path, "/")

	splitPart, err := strconv.Atoi(part[2])
	if err != nil{
		log.Println(err)
		return
	}
}



*/