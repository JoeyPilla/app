package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/JoeyPilla/rest-api-example/api"
	"github.com/JoeyPilla/rest-api-example/global"

	_ "github.com/lib/pq"
)

type Ready struct {
	Ready bool `json:"ready"`
}

type Motors struct {
	Motor1 string `json:"motor1"`
	Motor2 string `json:"motor2"`
	Motor3 string `json:"motor3"`
	Motor4 string `json:"motor4"`
	Motor5 string `json:"motor5"`
	Motor6 string `json:"motor6"`
	Motor7 string `json:"motor7"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Endpoint Hit")
	http.FileServer(http.Dir("<path to build>"))
}

func setMotors(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		resp := Motors{}
		json.NewDecoder(r.Body).Decode(&resp)
		global.MotorMap["motor1"], _ = strconv.Atoi(resp.Motor1)
		global.MotorMap["motor2"], _ = strconv.Atoi(resp.Motor2)
		global.MotorMap["motor3"], _ = strconv.Atoi(resp.Motor3)
		global.MotorMap["motor4"], _ = strconv.Atoi(resp.Motor4)
		global.MotorMap["motor5"], _ = strconv.Atoi(resp.Motor5)
		global.MotorMap["motor6"], _ = strconv.Atoi(resp.Motor6)
		global.MotorMap["motor7"], _ = strconv.Atoi(resp.Motor7)
	default:
		fmt.Fprintf(w, "Sorry, only POST methods are supported.")
	}
}

func getMotors(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		resp := Motors{}
		resp.Motor1 = strconv.Itoa(global.MotorMap["motor1"])
		resp.Motor2 = strconv.Itoa(global.MotorMap["motor2"])
		resp.Motor3 = strconv.Itoa(global.MotorMap["motor3"])
		resp.Motor4 = strconv.Itoa(global.MotorMap["motor4"])
		resp.Motor5 = strconv.Itoa(global.MotorMap["motor5"])
		resp.Motor6 = strconv.Itoa(global.MotorMap["motor6"])
		resp.Motor7 = strconv.Itoa(global.MotorMap["motor7"])
		json.NewEncoder(w).Encode(&resp)
	default:
		fmt.Fprintf(w, "Sorry, only GET methods are supported.")
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	for _, v := range global.MotorMap {
		if v != 0 {
			json.NewEncoder(w).Encode(Ready{Ready: true})
			return
		}
	}
	json.NewEncoder(w).Encode(Ready{Ready: false})
}

func handleRequests() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/setMotors", setMotors)
	http.HandleFunc("/getMotors", getMotors)
	api.AddRoutes()
	// raspberrypi.AddRoutes()
	log.Fatal(http.ListenAndServe(":8081", nil))

}

func main() {
	global.MotorMap = make(map[string]int)
	global.MotorMap["motor1"] = 0
	global.MotorMap["motor2"] = 0
	global.MotorMap["motor3"] = 0
	global.MotorMap["motor4"] = 0
	global.MotorMap["motor5"] = 0
	global.MotorMap["motor6"] = 0
	global.MotorMap["motor7"] = 0

	// raspberrypi.Initialize()

	db := api.Connect()
	handleRequests()
	defer db.Close()
}
