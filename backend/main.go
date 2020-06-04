package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JoeyPilla/rest-api-example/api"
	"github.com/JoeyPilla/rest-api-example/raspberrypi"

	_ "github.com/lib/pq"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	api.AddRoutes()
	raspberrypi.AddRoutes()
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func main() {
	api.Connect()
	handleRequests()
}
