package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"unicode"
)

var Artist []Artists

func GetData(data interface{}, url string) error {
	apiURL := url

	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error making GET request: ", err)
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}
	return nil
}

func AdditionalData(id int) error {
	location := StructLocations{}
	GetData(&location, "https://groupietrackers.herokuapp.com/api/locations/"+strconv.Itoa(id))

	Artist[id-1].Locations = location

	date := StructConcertDates{}
	GetData(&date, "https://groupietrackers.herokuapp.com/api/dates/"+strconv.Itoa(id))

	Artist[id-1].ConcertDates = date

	relation := StructRelations{}
	GetData(&relation, "https://groupietrackers.herokuapp.com/api/relation/"+strconv.Itoa(id))
	Artist[id-1].Relations = relation

	return nil
}

func IsValid(id string) bool {
	if id == "" {
		return false
	}
	for _, char := range id {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return id[0] != '0'
}

func IsRange(id int) bool {
	if id < 1 || id > 52 {
		return false
	}
	return true
}
