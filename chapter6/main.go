package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var spaceline = "SpaceX"
	var triptype bool

	var distance = 62100000 //км

	fmt.Printf("Spaceline Days Trip Type Price\n")
	for count := 0; count < 10; count++ {
		var k = rand.Intn(15)
		var shipspeed = k + 16                                   //км/с
		var flightduration = distance / shipspeed / 60 / 60 / 24 // days
		var price = k*1000000 + 36000000
		if triptype {
			price *= 2
		}
		fmt.Printf("%v %v %v %v", spaceline, flightduration, triptype, price)
		fmt.Println(count)
	}

}
