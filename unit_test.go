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

// этот тест проверяет, что наши обработчики правильно обрабатывают различные HTTP-методы
func TestHandlers_Methods(t *testing.T) {
	// Определение тестовых случаев: URL и соответствующий HTTP-метод
	testCases := []struct {
		url    string
		method string
	}{
		{"/", "GET"},
		{"/", "POST"},
		{"/", "PUT"},
		{"/", "DELETE"},
		{"/", "PATCH"},
		{"/artist/?id=1", "GET"},
		{"/artist/?id=1", "POST"},
		{"/artist/?id=1", "PUT"},
		{"/artist/?id=1", "DELETE"},
		{"/artist/?id=1", "PATCH"},
	}
	// Для каждого тестового случая...
	for _, tc := range testCases {
		req, err := http.NewRequest(tc.method, tc.url, nil)
		if err != nil {
			t.Fatal(err)
		}
		// Создание тестового recorder для записи ответа от обработчика
		rr := httptest.NewRecorder()

		// Определение обработчика, который будет использоваться в зависимости от URL-адреса
		var handler http.Handler
		if tc.url == "/" {
			handler = http.HandlerFunc(handlers.IndexPage)
		} else {
			handler = http.HandlerFunc(handlers.ArtistPage)
		}

		// Выполнение HTTP-запроса и запись ответа в recorder
		handler.ServeHTTP(rr, req)

		// Ожидаемый HTTP status code: http.StatusOK для GET, иначе http.StatusMethodNotAllowed
		expectedStatusCode := http.StatusMethodNotAllowed
		if tc.method == "GET" {
			expectedStatusCode = http.StatusOK
		}
		// Проверка соответствия фактического HTTP status code ожидаемому
		if status := rr.Code; status != expectedStatusCode {
			t.Errorf("%s request to %s returned wrong status code. Got %v, want %v", tc.method, tc.url, status, expectedStatusCode)
		}
	}
}

// этот тест проверяет, что indexHandler правильно обрабатывает запросы к несуществующим страницам
func TestIndexHandler_NotFound(t *testing.T) {
	// Создание HTTP-запроса для доступа к несуществующему URL-адресу
	req, err := http.NewRequest("GET", "/notfound", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Создание тестового recorder для записи ответа от обработчика
	rr := httptest.NewRecorder()
	// Выполнение HTTP-запроса и запись ответа в recorder
	handler := http.HandlerFunc(handlers.IndexPage)
	handler.ServeHTTP(rr, req)

	// Проверка, что фактический HTTP status code соответствует ожидаемому коду 404 (http.StatusNotFound)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Home handler returned wrong status code. Got %v, want %v", status, http.StatusNotFound)
	}
}

// этот тест проверяет, что artistHandler правильно обрабатывает запросы c неправильным id
func TestArtistPageHandler_InvalidID(t *testing.T) {
	// Создание HTTP-запроса с неправильным ID
	req, err := http.NewRequest("GET", "/artist/?id=invalid", nil)
	if err != nil {
		t.Fatal(err)
	}
	// Создание тестового recorder для записи ответа от обработчика
	rr := httptest.NewRecorder()

	// Выполнение HTTP-запроса и запись ответа в recorder
	handler := http.HandlerFunc(handlers.ArtistPage)
	handler.ServeHTTP(rr, req)

	// Проверка, что фактический HTTP status code соответствует ожидаемому коду 400 (http.StatusBadRequest)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("ArtistPage handler returned wrong status code. Got %v, want %v", status, http.StatusBadRequest)
	}
}
