package main

import (
	"fmt"
	"net/http"

	"github.com/Yandex-Practicum/go-rest-api-homework/cmd/handler"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()                       // Новый роутер
	handler := handler.CreateToDoListHandler() // Custom handler

	r.Get("/tasks", handler.GetAllTasks)            // Получение всех задач
	r.Post("/tasks", handler.AddTask)               // Добавление новой задачи
	r.Get("/tasks/{id}", handler.GetTaskByID)       // Получить задачу по ID
	r.Delete("/tasks/{id}", handler.DeleteTaskByID) // Удаляет задачу по ее ID

	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}
}
