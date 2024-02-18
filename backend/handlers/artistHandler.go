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

	// Извлекаем идентификатор артиста из параметра запроса.
	id := r.URL.Query().Get("id")

	// Проверяем, что идентификатор артиста содержит только цифры (здесь используется функция IsValid которую мы задаем в пакете utilities backend/utilities/utilities.go пройди туда, там объяснено как они работают)
	if !utilities.IsValid(id) {
		ErrorPage(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	if utilities.StartsWithZero(id) {
		ErrorPage(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	idd, err := strconv.Atoi(id)

	err = data.FetchDataFromJSON(&data.Artists, "https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		ErrorPage(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if !utilities.IsInRange(idd) {
		ErrorPage(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	err = data.GetDataForArtist(idd)
	if err != nil {
		ErrorPage(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	t, err := template.ParseFiles("frontend/html/artist.html")
	if err != nil {
		log.Println(err)
		ErrorPage(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	err = t.Execute(w, data.Artists[idd-1])
	if err != nil {
		http.Error(w, "Error executin file", http.StatusInternalServerError)
		return
	}
}
