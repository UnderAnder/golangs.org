package main

import "fmt"

type celsius float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

type fahrenheit float64

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

type sensor func(t celsius) (string, string)

func drawTable(header1, header2 string, tempMin, tempMax, step celsius, s sensor) {
	line := "===============\n"
	pattern := "|%6v|%6v|\n"
	fmt.Printf(line)
	fmt.Printf(pattern, header1, header2)
	rows := (tempMax - tempMin) / step
	for row := 0; row <= int(rows); row++ {
		cell1, cell2 := s(tempMin)
		fmt.Printf(pattern, cell1, cell2)
		tempMin += step
	}
	fmt.Printf(line)
}

func ctof(temp celsius) (string, string) {
	cell1 := fmt.Sprintf("%v", temp)
	cell2 := fmt.Sprintf("%v", temp.fahrenheit())
	return cell1, cell2
}

func ftoc(temp celsius) (string, string) {
	cell1 := fmt.Sprintf("%v", temp.fahrenheit())
	cell2 := fmt.Sprintf("%v", temp)
	return cell1, cell2
}

func main() {
	var tempMin, tempMax, step celsius = -40, 100, 5
	drawTable("°C", "°F", tempMin, tempMax, step, ctof)
	drawTable("°F", "°C", tempMin, tempMax, step, ftoc)
}
