package raspberrypi

import (
	"fmt"
	"os"
	"time"

	"github.com/stianeikeland/go-rpio"
)

var motors []rpio.Pin
var cupToMl float64 = 236.588
var ozToMl float64 = 29.5735

func Initialize() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	motors = []rpio.Pin{
		rpio.Pin(2),
		rpio.Pin(3),
		rpio.Pin(14),
		rpio.Pin(15),
		rpio.Pin(23),
		rpio.Pin(17),
		rpio.Pin(22),
	}

	for _, motor := range motors {
		motor.Output()
		motor.High()
	}

}

// motor1 0.667 ml/sec
// motor2 0.76 ml/sec
// motor3 0.5 ml/sec
// motor4 1.1 ml/sec
// motor5 1.0 ml/sec
// motor6 1.0 ml/sec
// motor7 1.0 ml/sec

func Pour(motor int, amount float64, unitOfMeasurement string, c chan float64) {
	if motor < 0 {
		c <- 0.0
		return
	}
	motors[motor].Low()
	pourRate := []float64{
		0.8,
		0.96,
		0.725,
		1.1,
		1.0,
		0.75,
		0.6,
	}

	t := float64(amount) / pourRate[motor]
	if unitOfMeasurement == "cup" || unitOfMeasurement == "cups" {
		t = t * cupToMl
	} else if unitOfMeasurement == "oz" {
		t = t * ozToMl
	}
	// write how long it will take
	fmt.Print(t)
	c <- t
	time.Sleep(time.Duration(t * float64(time.Second)))
	defer motors[motor].High()
}
