package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const hoursPerDay = 24
	var distance = rand.Intn(401000000-56000000+1) + 5600000 // km 96300000
	var speed = 100800                                       // km/h
	fmt.Printf("Distance to mars: %v km\nDays to Mars: %v", distance, distance/speed/hoursPerDay)
}
