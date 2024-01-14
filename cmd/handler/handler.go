package handler

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Yandex-Practicum/go-rest-api-homework/data"
	"github.com/go-chi/chi/v5"
)

type ToDoListHandler struct{}

func CreateToDoListHandler() *ToDoListHandler {
	return &ToDoListHandler{}
}

// GetAllTasks возвращает все имеющиеся задачи (эндпоинт: /tasks)
func (h *ToDoListHandler) GetAllTasks(w http.ResponseWriter, req *http.Request) {
	// Случай, когда нет задач в базе
	if len(data.Tasks) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Список задач пуст"))
		return
	}
	dataJson, err := json.Marshal(data.Tasks)
	if err != nil {
		http.Error(w, "error Marshal JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(dataJson))
}

// AddTask добавляет новую задачу (эндпоинт: /tasks)
func (h *ToDoListHandler) AddTask(w http.ResponseWriter, req *http.Request) {
	var buf bytes.Buffer
	_, err := buf.ReadFrom(req.Body) // Чтение JSON из тела запроса
	if err != nil {
		http.Error(w, "error read request", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()
	newTask := data.CreateTask() // Экзмепляр структуры Task
	err = json.Unmarshal(buf.Bytes(), newTask)
	if err != nil {
		http.Error(w, "error UnMarshal JSON", http.StatusBadRequest)
		return
	}
	// Если ID новой задачи равен ID уже существующей задачи
	if _, ok := data.Tasks[newTask.ID]; ok {
		http.Error(w, "error adding task", http.StatusBadRequest)
		return
	}
	data.Tasks[newTask.ID] = *newTask // Добавление новой задачи
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Задача добавлена"))
}

// GetTaskByID возвращает задачу по ее ID (эндпоинт: /tasks/{id})
func (h *ToDoListHandler) GetTaskByID(w http.ResponseWriter, req *http.Request) {
	idTask := chi.URLParam(req, "id")
	// Если ID искомой задачи нет в мапе
	if _, ok := data.Tasks[idTask]; !ok {
		http.Error(w, "Задачи нет в базе", http.StatusBadRequest)
		return
	}
	dataJson, err := json.Marshal(data.Tasks[idTask])
	if err != nil {
		http.Error(w, "error Marshal JSON", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(dataJson))
}

// DeleteTaskByID удаляет задачу по ее ID (эндпоинт: /tasks/{id})
func (h *ToDoListHandler) DeleteTaskByID(w http.ResponseWriter, req *http.Request) {
	idTask := chi.URLParam(req, "id")
	// Если ID искомой задачи нет в мапе
	if _, ok := data.Tasks[idTask]; !ok {
		http.Error(w, "Задачи нет в базе", http.StatusBadRequest)
		return
	}
	delete(data.Tasks, idTask)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Задача удалена"))
}
