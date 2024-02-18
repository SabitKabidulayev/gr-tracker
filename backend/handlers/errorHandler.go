package handlers

// в пакете handlers находятся handler-ы (обработчики). Handler представляет собой функцию, которая принимает HTTP запросы (methods) и возвращает HTTP ответы. Он используется для обработки запросов, поступающих на сервер, и выполнения необходимой логики в зависимости от характера запроса.

import (
	"fmt"
	"net/http"
	"text/template"
)

// ErrorStruct представляет структуру данных для передачи информации об ошибке на страницу ошибок.
type ErrorStruct struct {
	StatusCode string // Код ошибки.
	StatusText string // Текст ошибки.
}

// ErrorPage - обработчик для отображения страницы ошибок.
// w http.ResponseWriter: это интерфейс, который используется для записи ответа клиенту. С помощью него можно записывать данные, которые будут отправлены обратно клиенту, такие как HTML, JSON и т. д.
// r *http.Request: это объект, представляющий HTTP запрос, содержащий всю информацию о запросе, такую как метод, заголовки, параметры запроса и тело запроса.
func ErrorPage(w http.ResponseWriter, r *http.Request, statusCode int, statusText string) {
	// Создание экземпляра структуры ErrorStruct с переданными значениями statusCode и statusText.
	errorData := ErrorStruct{
		StatusCode: fmt.Sprint(statusCode), // Преобразование statusCode  в строку.
		StatusText: statusText,             // Задание текста статуса ошибки.
	}

	// Парсинг HTML шаблона страницы ошибок из файла "errpage.html".
	t, err := template.ParseFiles("./frontend/templates/errorPage.html")
	// Если произошла ошибка при парсинге шаблона, отправляем клиенту статус 500 "Внутренняя ошибка сервера".
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError))) // Отправка текста статуса ошибки в виде тела ответа
		return
	}

	// Установка HTTP статуса ошибки.
	w.WriteHeader(statusCode)

	// Вставка данных из структуры ErrorStruct в HTML шаблон и отправка результата клиенту.
	err = t.Execute(w, errorData)
	// Если произошла ошибка при выполнении шаблона, отправляем клиенту статус 500 "Внутренняя ошибка сервера".
	if err != nil {
		http.Error(w, "Error when executing", http.StatusInternalServerError)
		return
	}
}

// дальше переходи в backend/handlers/indexHandler.go
