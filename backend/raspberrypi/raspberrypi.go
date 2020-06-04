package raspberrypi

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio"
)

var (
	// Use mcu pin 10, corresponds to physical pin 19 on the pi
	pin = rpio.Pin(21)
)

func AddRoutes() {
	http.HandleFunc("/api/blink", blinkRoute)
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Set pin to output mode
	pin.Output()
	pin.High()
}

func blinkLight() {
	// Toggle pin 20 times
	for x := 0; x < 20; x++ {
		pin.Toggle()
		time.Sleep(time.Second / 5)
	}
}

func blinkRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Blink Endpoint Hit")
	go blinkLight()
}
