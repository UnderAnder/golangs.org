/*
Карликовая галактика в Большом Псе является ближайшей известной к Земле галактикой, что находится на расстоянии 236 000 000 000 000 000 км от нашего Солнца.
 Используйте константы для конвертации данного расстояния в световые годы.
*/

package main

import (
	"fmt"
)

func main() {
	const distance = 236000000000000000
	const lightSpeed = 299792
	const secondsPerDay = 86400
	const daysPerYear = 365

	const days = distance / lightSpeed / secondsPerDay / daysPerYear

	fmt.Println("Расстояние до галактики в Большом Псе составляет", days, "световых лет.")

}
