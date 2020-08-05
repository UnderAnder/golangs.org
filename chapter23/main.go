package main

import (
	"fmt"
	"math"
)

type world struct {
	radius float64
}

type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	Lat  float64 `json:"Latitude"`
	Long float64 `json:"Longitude"`
}

var (
	mars  = world{radius: 3389.5}
	earth = world{radius: 6371}
)

/* locations := []location{
{Name: "Columbia Memorial Station", Lat: -14.5684, Long: 175.472636},
{Name: "Challenger Memorial Station", Lat: -1.9462, Long: 354.4734},
{Name: "Bradbury Landing", Lat: -4.5895, Long: 137.4417},
} */

// вычисление расстояния через Сферическую теорему косинусов.
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.Lat))
	s2, c2 := math.Sincos(rad(p2.Lat))
	cLong := math.Cos(rad(p1.Long - p2.Long))
	return w.radius * math.Acos(s1*s2+c1*c2*cLong) // Использует поле радиуса world
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

// rad конвертирует градусы в радианы.
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

func main() {
	spirit := newLocation(coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'})
	opportunity := newLocation(coordinate{1, 56, 46.3, 'S'}, coordinate{354, 28, 24.2, 'E'})
	curiosity := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.12, 'E'})
	insight := newLocation(coordinate{4, 30, 0.0, 'N'}, coordinate{135, 54, 0, 'E'})

	dist := mars.distance(spirit, opportunity) // Использует метод distance для mars
	fmt.Printf("%.2f km\n", dist)              // Выводит: 9669.71 km

	fmt.Println("Spirit", spirit)
	fmt.Println("Opportunity", opportunity)
	fmt.Println("Curiosity", curiosity)
	fmt.Println("InSight", insight)
}
