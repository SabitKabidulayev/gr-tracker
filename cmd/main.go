package main

// в пакете main создается веб-сервер, который обрабатывает запросы к корневому URL, URL, начинающемуся с "/artist/", а также загружает статические файлы CSS из директории "frontend/css".

import (
	"groupie-tracker/backend/handlers"
	"log"
	"net/http"
)

func main() {
	// Создаем новый маршрутизатор
	// маршрутизатор (или роутер) является компонентом, который принимает входящие HTTP-запросы и определяет, какой обработчик (handler) должен быть вызван для каждого конкретного запроса на основе его пути (URL) и метода (method)
	mux := http.NewServeMux()

	// Устанавливаем обработчик для корневого URL "/"
	mux.HandleFunc("/", handlers.IndexPage)
	// Устанавливаем обработчик для URL, начинающегося с "/artist/"
	mux.HandleFunc("/artist/", handlers.ArtistPage)

	// Создаем файловый сервер для обслуживания статических файлов CSS
	fileServer := http.FileServer(http.Dir("frontend/css"))
	// Устанавливаем обработчик для URL-префикса "/static/", который перенаправляет запросы к статическим файлам CSS
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Выводим сообщение в лог о начале работы сервера
	log.Println("Starting server on http://localhost:8000")

	// Запускаем HTTP-сервер на порту 8080, используя созданный маршрутизатор mux для обработки запросов
	http.ListenAndServe(":8080", mux)
}
