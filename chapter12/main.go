package main

import (
	"fmt"
)

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"

	for i, c := range cipherText {
		if c >= 'A' && c <= 'Z' {
			if i < len(keyword) {
				c = c - rune(keyword[i]) // 'G'
				if c < 'A' {
					c += 26
				}
			} else {
				i -= len(keyword)
			}
		}
		fmt.Printf("%c", c)
	}

	for i := 0; i < len(cipherText); i++ {
		for k := 0; i < len(keyword); k++ {
			c := cipherText[i] + keyword[i]
			fmt.Printf("%c", c)
		}

	}
	fmt.Println('C' - 'G')
	fmt.Println('S' - 'O')
	fmt.Println('O' - 'L')
}
