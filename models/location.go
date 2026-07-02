package models

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type LocationResponse struct {
	Index []Location `json:"index"`
}
