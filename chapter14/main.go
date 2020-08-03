package main

import "fmt"

type celsius float64
type kelvin float64

// kelvinToCelsius converts °K to °C
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15) // Необходима конвертация типа
}

func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15) // Необходима конвертация типа
}

func main() {
	var c celsius = 127
	k := celsiusToKelvin(c)
	fmt.Print(c, "°C is", k, " °K")
}
