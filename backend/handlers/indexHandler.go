package handlers

// в пакете handlers находятся handler-ы (обработчики). Handler представляет собой функцию, которая принимает HTTP запросы (methods) и возвращает HTTP ответы. Он используется для обработки запросов, поступающих на сервер, и выполнения необходимой логики в зависимости от характера запроса.
// Основные виды HTTP запросов:
// GET: Запрос на получение данных с сервера
// POST: Запрос на отправку данных на сервер
// PUT: Запрос на обновление существующих данных на сервере
// DELETE: Запрос на удаление данных на сервере
// PATCH: Запрос на частичное обновление данных на сервере
// в нашем случае все indexHandler работает только с GET запросами (в случае других запосов будет выводится ошибка MethodNotAllowed)

import (
	"groupie-tracker/backend/data"
	"log"
	"net/http"
	"text/template"
)

// функция IndexPage - обработчик для главной страницы
// w http.ResponseWriter: это интерфейс, который используется для записи ответа клиенту. С помощью него можно записывать данные, которые будут отправлены обратно клиенту, такие как HTML, JSON и т. д.
// r *http.Request: это объект, представляющий HTTP запрос, содержащий всю информацию о запросе, такую как метод, заголовки, параметры запроса и тело запроса.
func IndexPage(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что URL запроса соответствует корневому пути "/"
	// Если URL не корневой, вызываем обработчик ошибки и возвращаем ошибку 404 "Страница не найдена"
	if r.URL.Path != "/" {
		ErrorPage(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	// Проверяем, что метод запроса - GET
	// Если метод запроса не GET, вызываем обработчик ошибки и возвращаем ошибку 405 "Метод не разрешен"
	if r.Method != http.MethodGet {
		ErrorPage(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	// Парсим HTML шаблон для домашней страницы
	t, err := template.ParseFiles("frontend/templates/indexPage.html")
	// Если произошла ошибка при парсинге шаблона, записываем ее в лог, вызываем обработчик ошибки и возвращаем ошибку 500 "Внутренняя ошибка сервера"
	if err != nil {
		log.Println(err)
		ErrorPage(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	// Получаем данные из удаленного JSON-файла (с помощью функции FetchDataFromJSON из пакетв data) и сохраняем их в структуру data.Artists
	err = data.FetchDataFromJSON(&data.Artists, "https://groupietrackers.herokuapp.com/api/artists")
	// Если произошла ошибка при получении данных, вызываем обработчик ошибки и возвращаем ошибку 500 "Внутренняя ошибка сервера"
	if err != nil {
		ErrorPage(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	// Вставляем данные из data.Artists в HTML шаблон и отправляем результат клиенту
	err = t.Execute(w, data.Artists)
	// Если произошла ошибка при выполнении шаблона, возвращаем ошибку 500 "Внутренняя ошибка сервера"
	if err != nil {
		http.Error(w, "Error when executing", http.StatusInternalServerError)
		return
	}
}

// дальше переходи в backend/handlers/artistHandler.go
