package main

import (
	"fmt"
	"net/http"
)

func index_handler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello, Go is neat!</h1>")
	fmt.Fprintf(w, "<p>Go is fast</p>")
	fmt.Fprintf(w, "<p>... and simple</p>")

}
// const usixteenbitmax float64 = 65535
// const kmh_multiple float64 = 1.60934

// type car struct {
// 	gas_pedal uint16
// 	brake_pedal uint16
// 	steering_wheel int16
// 	top_speed_kmh float64
// }
// // func about_handler (w http.ResponseWriter, r *http.Request) {
// // 	fmt.Fprintf(w, "That is the about page!")
// // }
// func (c car) kmh() float64 {
// 	return float64 (c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax)
// }
// func (c car) mph() float64 {
// 	return float64 (c.gas_pedal) * (c.top_speed_kmh/usixteenbitmax/kmh_multiple)
// }

// func (c *car) new_top_speed(newspeed float64) {
// 	c.top_speed_kmh = newspeed
// }

// func new_speed_car(c car, speed float64) car {
// 	c.top_speed_kmh = speed
// 	return c
// }
func main() {
	http.HandleFunc("/", index_handler)
	// http.HandleFunc("/about", about_handler)
	http.ListenAndServe(":8000", nil)
	// a_car := car{gas_pedal: 65000,
	// 							brake_pedal: 0,
	// 							steering_wheel: 12561,
	// 							top_speed_kmh: 225.0}
	
	// fmt.Println(a_car.gas_pedal)
	// fmt.Println("a car ",a_car.kmh())
	// fmt.Println("mile per hour", a_car.mph())
	// a_car = new_speed_car(a_car, 500)

	// fmt.Println("a car ",a_car.kmh())
	// fmt.Println("mile per hour", a_car.mph())
}
