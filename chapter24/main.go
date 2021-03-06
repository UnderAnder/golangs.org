package main

import (
	"fmt"
	"math"
)

type rover struct {
	gps
}

type gps struct {
	location   location
	destantion location
	world      world
}

type location struct {
	name      string
	lat, long float64
}

type world struct {
	radius float64
}

func (l location) description() string {
	return fmt.Sprint(l)
}

//реализуйте метод message, что возвращает строку с оставшимися километрами до пункта назначения.
func (g gps) message() string {
	return fmt.Sprint(g.world.distance(g.location, g.destantion))
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	cLong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*cLong) // Использует поле радиуса world
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func main() {
	var (
		mars = world{radius: 3389.5}
		// earth = world{radius: 6371}
	)

	bradbury := location{"Bradbury Landing", -4.5895, 137.4417}
	elysium := location{"Elysium Planitia", 4.5, 135.9}
	gps := gps{world: mars, location: bradbury, destantion: elysium}
	curiosity := rover{gps: gps}

	fmt.Printf("%+v\n", gps)
	fmt.Printf("км. до рункта назначения: %v\n", curiosity.message())
}
