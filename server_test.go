package test

// в пакете test мы пишем функции-тесты для проверки функций, которые мы использовали в проекте
// все наши тесты являются модульными тестами (Unit tests) - так как все они тестируют отдельные модули (функций) программы для проверки их работы
// также есть Нагрузочные тесты (Load tests) - тесты производительности, Стресс-тесты (Stress tests) - тест пределов выносливости программы, Интеграционные тесты (Integration tests) - тестирование взаимодействия между несколькими модулями и куча других. (могут спросить про это)

import (
	"encoding/json"
	"groupie-tracker/backend/data"
	"groupie-tracker/backend/handlers"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

// создаем структуру TestData с помощью которая будет использоваться в тесте для функции FetchDataFromJSON

type FetchTestData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// *testing.T - это объект, который предоставляет методы для работы с тестами. Этот объект передается в функции-тесты как аргумент, чтобы мы могли использовать его методы (t.Errorf, t.Run, t.Fatal) для выполнения проверок и сообщения о результатах теста

func TestFetchDataFromJSON(t *testing.T) {
	// Создаем тестовый HTTP-сервер с помощью httptest.NewServer, который будет возвращать заданные данные в формате JSON.
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Создаем тестовые данные mockData
		mockData := FetchTestData{Name: "Test Name", Email: "test@mail.com"}
		// Преобразуем тестовые данные в JSON.
		mockDataJSON, _ := json.Marshal(mockData)
		// Отправляем JSON-данные в ответ на запрос.
		w.Write(mockDataJSON)
	}))
	// После завершения теста закрываем тестовый сервер, чтобы освободить ресурсы.
	defer testServer.Close()

	// Создаем переменную для хранения данных, полученных из JSON.
	var testData FetchTestData
	// Получаем URL тестового сервера.
	url := testServer.URL
	// Ожидаемые данные, которые должны быть получены из JSON.
	expectedData := FetchTestData{Name: "Test Name", Email: "test@mail.com"}

	// Запускаем функцию FetchDataFromJSON для извлечения данных из JSON и сохранения их в переменную testData.
	err := data.FetchDataFromJSON(&testData, url)
	// Проверяем наличие ошибок в процессе извлечения данных.
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Сравниваем фактические данные с ожидаемыми.
	// Функция DeepEqual сравнивает два значения любого типа и возвращает true, если они структурно эквивалентны, то есть имеют одинаковую структуру и значения всех полей равны. Если хотя бы одно поле отличается, функция вернет false.
	if !reflect.DeepEqual(testData, expectedData) {
		t.Errorf("Unexpected data. Got %+v, expected %+v", testData, expectedData)
	}
}

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
