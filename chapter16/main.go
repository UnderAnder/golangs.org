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

func drawTable(header1, header2 string, rows celsius, s sensor) {
	line := "===============\n"
	fmt.Printf(line)
	fmt.Printf("| %4v | %4v |\n", header1, header2)
	for row := 0; row < int(rows); row++ {
		cell1, cell2 := s(rows)
		fmt.Printf("| %4v | %4v |\n", cell1, cell2)
	}
	fmt.Printf(line)
}

func ctof(temp celsius) (string, string) {
	cell1 := fmt.Sprintf("%v", temp)
	cell2 := fmt.Sprintf("%v", temp.fahrenheit())
	return cell1, cell2
}

func main() {
	var temp celsius = 40
	drawTable("°F", "°C", temp, ctof)

	/* 	for temp = 40; temp < 100; temp += 5 {
	   		fmt.Printf("| %v   | %v  |\n", temp, temp.fahrenheit())
	   	}
	   	for temp = 40; temp < 100; temp += 5 {
	   		fmt.Printf("| %v   | %v  |\n", temp.fahrenheit(), temp)
	   	} */
}
