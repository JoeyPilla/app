package main

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"time"

// 	"github.com/JoeyPilla/rest-api-example/api"
// 	"github.com/JoeyPilla/rest-api-example/ingredient"
// )

// type Drinks struct {
// 	Drinks []Ingredient `json:"drinks"`
// }

// type Ingredient struct {
// 	Ingredient string `json:"strIngredient1"`
// }

// func main() {
// 	db := api.Connect()
// 	defer db.Close()

// }

// func seedIngredients() {
// 	url := "https://www.thecocktaildb.com/api/json/v1/1/list.php?i=list"

// 	spaceClient := http.Client{
// 		Timeout: time.Second * 2, // Maximum of 2 secs
// 	}

// 	req, err := http.NewRequest(http.MethodGet, url, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	req.Header.Set("User-Agent", "spacecount-tutorial")

// 	res, getErr := spaceClient.Do(req)
// 	if getErr != nil {
// 		log.Fatal(getErr)
// 	}

// 	body, readErr := ioutil.ReadAll(res.Body)
// 	if readErr != nil {
// 		log.Fatal(readErr)
// 	}

// 	drinks := Drinks{}
// 	jsonErr := json.Unmarshal(body, &drinks)
// 	if jsonErr != nil {
// 		log.Fatal(jsonErr)
// 	}

// 	for _, v := range drinks.Drinks {
// 		ingredient.AddIngredient(v.Ingredient)
// 	}
// }
