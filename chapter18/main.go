package main

import (
	"fmt"
	"sort"
)

// terraform ни к чему не приводит
func terraform(planets []string) {
	for i := range planets {
		planets[i] = "Новый " + planets[i]
	}
}

func main() {
	planets := []string{
		"Меркурий", "Венера", "Земля", "Марс",
		"Юпитер", "Сатурн", "Уран", "Нептун",
	}

	//sort.Strings(planets)
	terraform(planets[3:7])
	sort.StringSlice(planets).Sort() // Сортирует planets в алфавитном порядке
	fmt.Println(planets)             // Выводит: [Венера Земля Марс Меркурий Нептун Сатурн Уран Юпитер]
}
