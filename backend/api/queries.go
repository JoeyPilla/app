package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Todo struct {
	Name      string
	Completed bool
	Due       time.Time
}

type Todos []Todo

func handleApi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Api Endpoint Hit")
}

func AddRoutes() {
	http.HandleFunc("/api", handleApi)
	http.HandleFunc("/api/todo/all", allTodos)
}

func allTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}
	json.NewEncoder(w).Encode(todos)
}
