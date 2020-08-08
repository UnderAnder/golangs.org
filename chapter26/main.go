package main

import (
	"fmt"
	"math/rand"
	"time"
)

type sheep struct {
	name string
}
type cow struct {
	name string
}

type animal interface {
	walk() string
	eat() string
}

func (a sheep) String() string {
	return a.name
}

func (a cow) String() string {
	return a.name
}
func (a sheep) walk() string {
	return fmt.Sprint("Идет прямо в гору")
}

func (a cow) walk() string {
	return fmt.Sprint("Бредет по лугу")
}

func (a sheep) eat() string {
	food := [...]string{"мороженное", "бананы", "хлеб", "морковь"}
	f := food[rand.Intn(len(food))]
	return fmt.Sprintf("%v", f)
}

func (a cow) eat() string {
	switch rand.Intn(3) {
	case 0:
		return "морковку"
	case 1:
		return "салат-латук"
	default:
		return "редиску"
	}
}

func step(a animal) {
	switch rand.Intn(2) {
	case 0:
		fmt.Printf("%v %v.\n", a, a.walk())
	default:
		fmt.Printf("%v ест %v.\n", a, a.eat())
	}
}

const sunrise, sunset = 8, 18

func main() {
	rand.Seed(time.Now().UnixNano())
	
	s := []animal{sheep{name: "Dolly"}, cow{name: "Буренка"}}

	var sol, hour int

	for {
		fmt.Printf("%2d:00 ", hour)
		if hour < sunrise || hour >= sunset {
			fmt.Println("Животные спят.")
		} else {
			i := rand.Intn(len(s))
			step(s[i])
		}

		time.Sleep(500 * time.Millisecond)

		hour++
		if hour >= 24 {
			hour = 0
			sol++
			if sol >= 3 {
				break
			}
		}
	}
}
