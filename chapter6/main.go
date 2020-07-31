package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var spaceline = [3]string{"SpaceX", "Space Adventures", "Virgin Galactic"}
	var distance = 62100000 //км

	fmt.Printf("Spaceline         Days Trip Type Price\n")
	for count := 0; count < 10; count++ {
		var triptype string
		var k = rand.Intn(15)
		var shipspeed = k + 16                                   //км/с
		var flightduration = distance / shipspeed / 60 / 60 / 24 // days
		var price = k*1000000 + 36000000
		if rand.Intn(2) == 1 {
			triptype = "Round-trip"
			price = price * 2
		} else {
			triptype = "One-way"
		}

		fmt.Printf("%-16v %4v %-10v $%4v\n", spaceline[rand.Intn(3)], flightduration, triptype, price)
	}

}
