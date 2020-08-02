/*
Представьте, что вам нужно накопить денег на подарок другу.
Напишите программу, которая случайным образом размещает монеты пять ($0.05), десять ($0.10) и двадцать пять ($0.25) центов в пустую копилку до тех пор,
 пока внутри не будет хотя бы двадцать долларов ($20.00).
 Пускай после каждого пополнения копилки текущий баланс отображается на экране, отформатированный с нужной шириной и точностью.
*/
package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	piggyBank := 0.0

	for piggyBank < 20.00 {
		switch rand.Intn(3) {
		case 0:
			piggyBank += 0.05
		case 1:
			piggyBank += 0.10
		case 2:
			piggyBank += 0.25
		}
		fmt.Printf("В банке $%5.2f\n", piggyBank)
	}
}