/* Напишите программу, что конвертирует строки в булевы значения:

   Строки «true», «yes» или «1» соответствуют значению true;
   Строки «false», «no» или «0» соответствуют значению false;
   Для других значений выводит сообщение об ошибке.
*/
package main

import "fmt"

func main() {
	yesNo := "1"

	var launch bool

	switch yesNo {
	case "true", "yes", "1":
		launch = true
	case "false", "no", "0":
		launch = false
	default:
		fmt.Println(yesNo, "is not valid")
	}

	fmt.Println("Ready for launch:", launch) // Выводит: Ready for launch: true
}
