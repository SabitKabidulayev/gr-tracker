package data

import (
	"encoding/json"
	"groupie-tracker/backend/models"
	"net/http"
)

const (
	artistsApiURL   = "https://groupietrackers.herokuapp.com/api/artists"
	locationsApiURL = "https://groupietrackers.herokuapp.com/api/locations"
	datesApiURL     = "https://groupietrackers.herokuapp.com/api/dates"
	relationApiURL  = "https://groupietrackers.herokuapp.com/api/relation"
)

func GetData() (models.AllArtists, error) {
	art, err := http.Get(artistsApiURL)
	if err != nil {
		return nil, err
	}
	defer art.Body.Close()

	lctn, err := http.Get(locationsApiURL)
	if err != nil {
		return nil, err
	}
	defer lctn.Body.Close()

	date, err := http.Get(datesApiURL)
	if err != nil {
		return nil, err
	}
	defer date.Body.Close()

	rltn, err := http.Get(relationApiURL)
	if err != nil {
		return nil, err
	}
	defer rltn.Body.Close()

	var artists models.AllArtists

	err = json.NewDecoder(art.Body).Decode(&artists)
	if err != nil {
		return artists, err
	}

	var (
		locationDates models.IndexRelations
		location      models.IndexLocations
		dates         models.IndexConcerts
	)

	err = json.NewDecoder(date.Body).Decode(&dates)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(lctn.Body).Decode(&location)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(rltn.Body).Decode(&locationDates)
	if err != nil {
		return nil, err
	}
	for i := range artists {
		artists[i].ConcertDates = dates.ConcertDates[i]
		artists[i].Locations = location.Locations[i]
		artists[i].Relations = locationDates.Relations[i]
	}
	return artists, nil
}
