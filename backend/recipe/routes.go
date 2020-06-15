package recipe

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func AddRoutes() {
	http.HandleFunc("/api/recipe/all", allRecipes)
	http.HandleFunc("/api/recipe", recipe)
	http.HandleFunc("/api/recipe/available", availableRecipes)
}

func allRecipes(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getAllRecipes())
}

func availableRecipes(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getAvailableRecipes())
}

func RecipeById(w http.ResponseWriter, r *http.Request, id int) {

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getRecipeById(id))
}

func RecipeByName(w http.ResponseWriter, r *http.Request, name string) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getRecipeByName(name))
}

func recipe(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			name := r.URL.Query().Get("name")
			RecipeByName(w, r, name)
			return
		}
		RecipeById(w, r, id)
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		name := r.FormValue("name")
		json.NewEncoder(w).Encode(addRecipe(name))
	case "DELETE":
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			fmt.Fprintf(w, "Sorry, error deleting.")
			return
		}
		count := deleteRecipe(id)
		fmt.Fprintf(w, fmt.Sprintf("%d records deleted.", count))
	default:
		fmt.Fprintf(w, "Sorry, only GET, POST, and DELETE methods are supported.")
	}
}
