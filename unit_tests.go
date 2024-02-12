package test

import (
	"groupie-tracker/backend/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.IndexPage)

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
	handler := http.HandlerFunc(handlers.IndexPage)

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
	handler := http.HandlerFunc(handlers.IndexPage)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("Home handler returned wrong status code. Got %v, want %v", status, http.StatusMethodNotAllowed)
	}
}

func TestArtistPageHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/artist?id=1", nil)
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
	req, err := http.NewRequest("GET", "/artist?id=invalid", nil)
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

// type TestData struct {
// 	Name  string `json:"name"`
// 	Email string `json:"email"`
// }

// func TestGetData(t *testing.T) {
// 	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		mockData := TestData{Name: "John Doe", Email: "john@example.com"}
// 		mockBody, _ := json.Marshal(mockData)
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		w.Write(mockBody)
// 	}))nc TestGetData(t *testing.T) {
// 	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		mockData := TestData{Name: "John Doe", Email: "john@example.com"}
// 		mockBody, _ := json.Marshal(mockData)
// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusOK)
// 		w.Write(mockBody)
// 	}))
// 	defer testServer.Close()

// 	var data TestData
// 	url := testServer.URL
// 	expectedData := TestData{Name: "John Doe", Email: "john@example.com"}

// 	err := handlers.GetData(&data, url)
// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
// 	}

// 	if !reflect.DeepEqual(data, expectedData) {
// 		t.Errorf("Unexpect
// 	defer testServer.Close()

// 	var data TestData
// 	url := testServer.URL
// 	expectedData := TestData{Name: "John Doe", Email: "john@example.com"}

// 	err := handlers.GetData(&data, url)
// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
// 	}

// 	if !reflect.DeepEqual(data, expectedData) {
// 		t.Errorf("Unexpected data. Got %+v, expected %+v", data, expectedData)
// 	}
// }
