package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

// sensor function type
type sensor func() kelvin

func realSensor() kelvin {
	return 0 // TODO: имплементировать реальный сенсор
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin { // Объявляет и возвращает анонимную функцию
		return s() + offset
	}
}

func main() {
	var v kelvin = 4
	sensor := calibrate(fakeSensor, v)
	fmt.Println(sensor())
	fmt.Println(sensor())
	fmt.Println(sensor())
}
