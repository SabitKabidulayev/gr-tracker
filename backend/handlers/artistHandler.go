package handlers

import (
	"groupie-tracker/backend/data"
	"groupie-tracker/backend/utilities"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// ArtistPage обрабатывает запросы для страницы  с данныси о группе
func ArtistPage(w http.ResponseWriter, r *http.Request) {
	// Проверяем, что метод запроса GET.
	// Если метод запроса не GET, вызываем обработчик ошибки и возвращаем ошибку 405 "Метод не разрешен"
	if r.Method != http.MethodGet {
		ErrorPage(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	// Проверяем, что путь запроса соответствует /artist/
	// Если URL не соответствует /artist/, вызываем обработчик ошибки и возвращаем ошибку 404 "Страница не найдена"
	if r.URL.Path != "/artist/" {
		ErrorPage(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	// Извлекаем id артиста из параметра запроса.
	idString := r.URL.Query().Get("id")

	// Проверяем, что id содержит только цифры (здесь используется функция IsValid и StartsWithZero которую мы задаем в пакете utilities backend/utilities/utilities.go пройди туда, там объяснено как они работают)
	// Иначе возвращаем ошибку 400 "Bad request"  с помощью обработчика ошибки
	if !utilities.IsValid(idString) {
		ErrorPage(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	// Проверяем, что id не начинается с нуля.
	// Иначе возвращаем ошибку 400 "Bad request" с помощью обработчика ошибки
	if utilities.StartsWithZero(idString) {
		ErrorPage(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	// Преобразуем id в число.
	idInt, err := strconv.Atoi(idString)
	// Если произошла ошибка при преобразовании возвращаем ошибку 400 "Bad request"  с помощью обработчика ошибки
	if err != nil {
		ErrorPage(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	// Получаем данные из удаленного JSON-файла (с помощью функции FetchDataFromJSON из пакетв data) и сохраняем их в структуру data.Artists
	err = data.FetchDataFromJSON(&data.Artists, "https://groupietrackers.herokuapp.com/api/artists")
	// Если произошла ошибка при получении данных, вызываем обработчик ошибки и возвращаем ошибку 500 "Внутренняя ошибка сервера"
	if err != nil {
		ErrorPage(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	// Проверяем, что id группы находится в допустимом диапазоне. (функция IsInRange которую мы задаем в пакете utilities backend/utilities/utilities.go)
	if !utilities.IsInRange(idInt) {
		ErrorPage(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	// Получаем данные о группе по ее id. (функция GetDataForArtist из пакета data)
	err = data.GetDataForArtist(idInt)
	// Если произошла ошибка при получении данных, вызываем обработчик ошибки и возвращаем ошибку 500 "Внутренняя ошибка сервера"
	if err != nil {
		ErrorPage(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	// Парсим HTML шаблон для страницы группы
	t, err := template.ParseFiles("frontend/html/artist.html")
	// Если произошла ошибка при парсинге шаблона, записываем ее в лог, вызываем обработчик ошибки и возвращаем ошибку 500 "Внутренняя ошибка сервера"
	if err != nil {
		log.Println(err)
		ErrorPage(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	// Вставляем данные из data.Artists[idInt-1] в HTML шаблон и отправляем результат клиенту. Поскольку массивы индексируются с нуля, нам нужно вычесть 1 из idInt, чтобы получить правильный индекс в массиве.)
	err = t.Execute(w, data.Artists[idInt-1])
	// Если произошла ошибка при выполнении шаблона, возвращаем ошибку 500 "Внутренняя ошибка сервера"
	if err != nil {
		http.Error(w, "Error executin file", http.StatusInternalServerError)
		return
	}
}

// последняя часть backend это тесты, переходи в unit_test.go
