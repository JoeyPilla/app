package api

import (
	"fmt"
	"net/http"

	"github.com/JoeyPilla/rest-api-example/drink"
	"github.com/JoeyPilla/rest-api-example/ingredient"
	"github.com/JoeyPilla/rest-api-example/recipe"
)

func handleApi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Api Endpoint Hit")
}

func AddRoutes() {
	http.HandleFunc("/api", handleApi)
	ingredient.AddRoutes()
	recipe.AddRoutes()
	drink.AddRoutes()
}
