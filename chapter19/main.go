package main

import "fmt"

// dump для длины, вместимости и содержимого среза
func dump(label int, slice []string) {
	fmt.Printf("%v: длина %v, вместимость %v\n", label, len(slice), cap(slice))
}

func main() {
	planets := make([]string, 10)
	lastCap := cap(planets)
	for i := 0; i < 100; i++ {
		planets = append(planets, "New planet")
		if cap(planets) != lastCap {
			dump(i, planets)
			lastCap = cap(planets)
		}

	}
}
