package drink

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/JoeyPilla/rest-api-example/global"
	"github.com/JoeyPilla/rest-api-example/raspberrypi"
)

func AddRoutes() {
	http.HandleFunc("/api/drink/available", AvailableDrinks)
	http.HandleFunc("/api/drink/all", allDrinks)
	http.HandleFunc("/api/drink/pour", pourDrink)
	http.HandleFunc("/api/drink", drink)
}

func allDrinks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getAllDrinks())
}

func AvailableDrinks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getAvailableDrinks())
}

func DrinkByIngredient(w http.ResponseWriter, r *http.Request, id int) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getDrinkByIngredient(id))
}

func DrinkByRecipe(w http.ResponseWriter, r *http.Request, id int) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(getDrinkByRecipe(id))
}

func drink(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id, err := strconv.Atoi(r.URL.Query().Get("ingredientId"))
		if err == nil {
			DrinkByIngredient(w, r, id)
			return
		}
		id, err = strconv.Atoi(r.URL.Query().Get("recipeId"))
		if err == nil {
			DrinkByRecipe(w, r, id)
			return
		}
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		ingredientId, _ := strconv.Atoi(r.FormValue("ingredientId"))
		recipeId, _ := strconv.Atoi(r.FormValue("recipeId"))
		measure, _ := strconv.ParseFloat(r.FormValue("measure"), 64)
		unitOfMeasurement := r.FormValue("unitOfMeasurement")
		count := addDrink(ingredientId, recipeId, measure, unitOfMeasurement)
		fmt.Fprintf(w, fmt.Sprintf("%d records inserted.", count))
	case "DELETE":
		ingredientId, err := strconv.Atoi(r.URL.Query().Get("ingredientId"))
		recipeId, err := strconv.Atoi(r.URL.Query().Get("recipeId"))
		if err != nil || ingredientId < 1 || recipeId < 1 {
			fmt.Fprintf(w, "Sorry, error deleting.")
			return
		}
		count := deleteDrink(ingredientId, recipeId)
		fmt.Fprintf(w, fmt.Sprintf("%d records deleted.", count))
	default:
		fmt.Fprintf(w, "Sorry, only GET, POST, and DELETE methods are supported.")
	}
}

func motorToInt(motor string) (int, error) {
	return strconv.Atoi(motor[len(motor)-1:])
}

func getMotor(id string) int {
	intId, _ := strconv.Atoi(id)
	for k, v := range global.MotorMap {
		if v == intId {
			motorNumber, err := motorToInt(k)
			if err != nil {
				return -1
			}
			return motorNumber - 1
		}
	}
	return -1
}

func pourDrink(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		w.Header().Add("Content-Type", "application/json")
		id, err := strconv.Atoi(r.URL.Query().Get("recipeId"))
		if err == nil {
			c := make(chan float64)
			var time float64 = 0.0
			for _, ingredient := range getDrinkByRecipe(id) {
				go raspberrypi.Pour(getMotor(ingredient.Ingredient.ID), ingredient.Measure, ingredient.UnitOfMeasurement, c)
				t := <-c
				if time < t {
					time = t
				}
			}
			json.NewEncoder(w).Encode(time)
		}
	default:
		fmt.Fprintf(w, "Sorry, only GET methods are supported.")
	}
}
