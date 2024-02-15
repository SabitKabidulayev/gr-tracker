package main

import (
	"encoding/json"
	"groupie-tracker/backend/data"
	"groupie-tracker/backend/handlers"
	"groupie-tracker/backend/utilities"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.Home)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Home handler returned wrong status code. Got %v, want %v", status, http.StatusOK)
	}
}

func TestHomeHandler_NotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/notfound", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.Home)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Home handler returned wrong status code. Got %v, want %v", status, http.StatusNotFound)
	}
}

func TestHomeHandler_MethodNotAllowed(t *testing.T) {
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.Home)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("Home handler returned wrong status code. Got %v, want %v", status, http.StatusMethodNotAllowed)
	}
}

func TestArtistPageHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/artist/?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.ArtistPage)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("ArtistPage handler returned wrong status code. Got %v, want %v", status, http.StatusOK)
	}
}

func TestArtistPageHandler_InvalidID(t *testing.T) {
	req, err := http.NewRequest("GET", "/artist/?id=invalid", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.ArtistPage)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("ArtistPage handler returned wrong status code. Got %v, want %v", status, http.StatusBadRequest)
	}
}

type TestData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func TestFetchData(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mockData := TestData{Name: "John Doe", Email: "john@example.com"}
		mockBody, _ := json.Marshal(mockData)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(mockBody)
	}))
	defer testServer.Close()

	var testData TestData
	url := testServer.URL
	expectedData := TestData{Name: "John Doe", Email: "john@example.com"}

	err := data.FetchDataFromJSON(&testData, url)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(testData, expectedData) {
		t.Errorf("Unexpected data. Got %+v, expected %+v", testData, expectedData)
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		id       string
		expected bool
	}{
		{"123", true},
		{"asc", false},
	}

	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			result := utilities.IsValid(tt.id)
			if result != tt.expected {
				t.Errorf("IsValid(%s) = %v, expected %v", tt.id, result, tt.expected)
			}
		})
	}
}

func TestIsRange(t *testing.T) {
	tests := []struct {
		id       int
		expected bool
	}{
		{1, true},
		{0, false},
		{53, false},
	}

	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.id), func(t *testing.T) {
			result := utilities.IsRange(tt.id)
			if result != tt.expected {
				t.Errorf("IsRange(%d) = %v, expected %v", tt.id, result, tt.expected)
			}
		})
	}
}
