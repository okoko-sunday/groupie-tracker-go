package api

import (
	"encoding/json"
	"fmt"
	"groupie/models"
	"net/http"
)


const (
	artistsURL = "https://groupietrackers.herokuapp.com/api/artists"
	locationURL = "https://groupietrackers.herokuapp.com/api/locations"
	datesURL = "https://groupietrackers.herokuapp.com/api/dates"
	relationURL = "https://groupietrackers.herokuapp.com/api/relation"
)

func FetchArtists()([]models.Artist, error){

	resp, err := http.Get(artistsURL)
	if err != nil{
		return nil, err
	}

	defer resp.Body.Close()


	if resp.StatusCode != http.StatusOK{
		return nil, fmt.Errorf("Unexpected status code: %s", resp.Status)
	}

	var artists []models.Artist

	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil{
		return nil, err
	}

	return artists, nil
}



func FetchLocation()(models.LocationResponse, error){
	
	resp, err := http.Get(locationURL)
	if err != nil{
		return models.LocationResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		return models.LocationResponse{}, fmt.Errorf("Unexpected status code: %s", resp.Status)
	}

	var locations models.LocationResponse

	err = json.NewDecoder(resp.Body).Decode(&locations)
	if err != nil{
		return models.LocationResponse{}, err
	}

	return locations, nil

}


func FetchDate()(models.DateResponse, error){
	
	resp, err := http.Get(datesURL)
	if err != nil{
		return models.DateResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		return models.DateResponse{}, fmt.Errorf("Unexpected status code: %s", resp.Status)
	}

	var dates models.DateResponse

	err = json.NewDecoder(resp.Body).Decode(&dates)
	if err != nil{
		return models.DateResponse{}, err
	}

	return dates, nil

}

func FetchRelation()(models.RelationResponse, error){
	
	resp, err := http.Get(relationURL)
	if err != nil{
		return models.RelationResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		return models.RelationResponse{}, fmt.Errorf("Unexpected status code: %s", resp.Status)
	}

	var relations models.RelationResponse

	err = json.NewDecoder(resp.Body).Decode(&relations)
	if err != nil{
		return models.RelationResponse{}, err
	}

	return relations, nil

}

/*
//1  the function name will change which is FetchArtists
//2 the input parameter will chnage which is []models.Artist
func FetchArtists() ([]models.Artist, error) {


	//3 the http.Get will change also
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("unexpected response status code" + resp.Status)
	}

	//4 the variable name and the already established variable type will change
	var artists []models.Artist

	//5 the &artists will also change
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		return nil, err
	}

	//6 the return value will change
	return artists, nil
}


// that means in a function 6 things will change

*/