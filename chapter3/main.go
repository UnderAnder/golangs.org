package main

import "fmt"

func main() {
	const hoursPerDay = 24
	var distance = 96300000                                    // km
	var speed = 100800                                         // km/h
	fmt.Printf("Days to Mars: %v", distance/speed/hoursPerDay) // 39
}
