package main

import (
	"fmt"
	"sort"
)

type person struct {
	age int
}

func (p *person) birthday() {
	//p.age++ // разыменование указателя nil
	if p == nil {
		return
	}
	p.age++
}

func sortStrings(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool { return s[i] < s[j] }
	}
	sort.Slice(s, less)
}

func sortLen(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool { return len(s[i]) < len(s[j]) }
	}
	sort.Slice(s, less)
}

func mirepoix(ingredients []string) []string {
	return append(ingredients, "onion", "carrot", "celery")
}

func main() {
	var nowhere *int
	//fmt.Println(nowhere)  // Выводит: <nil>
	//fmt.Println(*nowhere) // Выводит: nil pointer dereference
	if nowhere != nil {
		fmt.Println(*nowhere)
	}

	var nobody *person
	fmt.Println(nobody) // Выводит: <nil>

	nobody.birthday()

	//////////////////////////////////////////////

	var fn func(a, b int) int
	fmt.Println(fn == nil) // Выводит: true

	food := []string{"onion", "carrot", "celery", "рис", "чай", "квас", "rice", "apple"}
	sortStrings(food, nil)
	fmt.Println(food) // Выводит: [carrot celery onion]
	//sortLen(food, nil)
	//fmt.Println(food)
	sortStrings(food, func(i, j int) bool { return len(food[i]) < len(food[j]) })
	fmt.Println(food)

	//////////////////////////////////////////////

	var soup []string
	fmt.Println(soup == nil) // Выводит: true

	for _, ingredient := range soup {
		fmt.Println(ingredient)
	}

	fmt.Println(len(soup)) // Выводит: 0

	soup = append(soup, "onion", "carrot", "celery")
	fmt.Println(soup) // Выводит: [onion carrot celery]

	soup2 := mirepoix(nil)
	fmt.Println(soup2) // Выводит: [onion carrot celery]

	soup3 := mirepoix(food)
	fmt.Println(soup3)

	///////////////////////////////////////////////////////

	var soup4 map[string]int
	fmt.Println(soup == nil) // Выводит: true

	measurement, ok := soup4["onion"]
	if ok {
		fmt.Println(measurement)
	}

	for ingredient, measurement := range soup4 {
		fmt.Println(ingredient, measurement)
	}

	for ingredient, measurement := range soup3 {
		fmt.Println(ingredient, measurement)
	}

	//////////////////////////////////////////
	var v interface{}
	fmt.Printf("%T %v %v\n", v, v, v == nil) // Выводит: <nil> <nil> true

	var p *int
	v = p
	//Во избежание неприятных сюрпризов при сравнении интерфейсов с nil лучше явно использовать идентификатор nil
	// вместо указания на переменную, что содержит nil.
	fmt.Printf("%T %v %v\n", v, v, v == nil) // Выводит: *int <nil> false
	fmt.Printf("%#v\n", v) // Выводит: (*int)(nil)

}
