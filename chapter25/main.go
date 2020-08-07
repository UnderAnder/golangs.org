package main

import (
	"fmt"
)

type coordinate struct {
	d, m, s float64
	h       rune
}

// String форматирует координаты DMS
func (c coordinate) String() string {
	return fmt.Sprintf("%v°%v'%.1f\" %c", c.d, c.m, c.s, c.h)
}

// location с широтой и долготой в десятичных градусах
type location struct {
	lat, long coordinate
}

// String форматирует location с широтой и долготой
func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

// decimal конвертирует координаты d/m/s в десятичные градусы.
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func main() {
	elysium := location{
		lat:  coordinate{4, 30, 0.0, 'N'},
		long: coordinate{135, 54, 0.0, 'E'},
	}
	fmt.Println("Elysium Planitia is at", elysium) // Выводит: Elysium Planitia is at 4°30’0.0” N, 135°54’0.0” E
}
