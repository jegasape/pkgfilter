package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Todos struct {
	Id        int    `json:"Id" `
	Title     string `json:"Title"`
	Completed bool   `json:"Completed"`
}

type TodosFunc interface {
	getTodo(w http.ResponseWriter, r *http.Request)
}

type TodoService struct {
	mu    sync.Mutex
	todos []Todos
}

func (t *TodoService) getTodo(w http.ResponseWriter, r *http.Request) {
	t.mu.Lock()
	defer t.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t.todos)
}

func main() {
	service := &TodoService{
		todos: []Todos{
			{Id: 1, Title: "Homework", Completed: true},
		},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			service.getTodo(w, r)
		}

	})

	fmt.Println("Starting server on localhost:8080")

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

