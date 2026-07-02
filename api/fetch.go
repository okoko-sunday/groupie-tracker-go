package api

import (
	"encoding/json"
	"errors"
	"groupie/models"
	"net/http"
)


func FetchArtists()([]models.Artist, error){

	// here i read from the Api
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	// this checks for error  
	if err != nil{
		return nil, err
	}

	// this helps close the resource when the function closes
	defer resp.Body.Close()


	if resp.StatusCode != http.StatusOK{
		return nil, errors.New("unexpected response status code"+ resp.Status)
	}

	// a variable to store the json 
	var artists []models.Artist



	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil{
		return nil, err
	}
	
	
	return artists, nil
}