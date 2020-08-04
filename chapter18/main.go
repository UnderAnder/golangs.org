package main

import (
	"fmt"
)

type planets []string

func (c planets) terraform() {
	for i := range c {
		c[i] = "Новый " + c[i]
	}
}

func main() {
	solsys := planets{
		"Меркурий", "Венера", "Земля", "Марс",
		"Юпитер", "Сатурн", "Уран", "Нептун",
	}

	//sort.Strings(planets)
	//sort.StringSlice(solsys).Sort() // Сортирует planets в алфавитном порядке
	solsys[3:4].terraform()
	solsys[6:].terraform()
	fmt.Println(solsys) // Выводит: [Венера Земля Марс Меркурий Нептун Сатурн Уран Юпитер]
}
