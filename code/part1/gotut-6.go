package main

import "fmt"

type car struct {
	gas_pedal      uint16 //min: 0,      max: 65535
	brake_pedal    uint16 //min: 0,      max: 65535
	steering_wheel int16  //min: -32768  max: 32768
	top_speed_kmh  float64
}

func main() {
	a_car := car{gas_pedal: 22341, brake_pedal: 0, steering_wheel: 12562, top_speed_kmh: 225.0}
	// a_car := car{22314, 0, 12562, 225.0}
	fmt.Println(a_car.gas_pedal)

}
