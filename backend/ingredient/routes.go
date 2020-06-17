package ingredient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func AddRoutes() {
	http.HandleFunc("/api/ingredient/all", allIngredients)
	http.HandleFunc("/api/ingredient", ingredient)
}

func allIngredients(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getAllIngredients())
}

func IngredientById(w http.ResponseWriter, r *http.Request, id int) {

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getIngredientById(id))
}

func IngredientByName(w http.ResponseWriter, r *http.Request, name string) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getIngredientByName(name))
}

func ingredient(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			name := r.URL.Query().Get("name")
			IngredientByName(w, r, name)
			return
		}
		IngredientById(w, r, id)
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		name := r.FormValue("name")
		count := AddIngredient(name)
		fmt.Fprintf(w, fmt.Sprintf("%d records inserted.", count))
	case "DELETE":
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			fmt.Fprintf(w, "Sorry, error deleting.")
			return
		}
		count := deleteIngredient(id)
		fmt.Fprintf(w, fmt.Sprintf("%d records deleted.", count))
	default:
		fmt.Fprintf(w, "Sorry, only GET, POST, and DELETE methods are supported.")
	}
}
