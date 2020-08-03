package main

import (
	"fmt"
)

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	word := ""

	/* 	for i, c := range cipherText {
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
	*/
	for i := 0; i < len(cipherText); i++ {
		// 'A' = 1, 'Z' = 26

		c := cipherText[i%len(cipherText)] - (keyword[i%len(keyword)] - 'A')
		if c < 'A' {
			c += 26
		}

		fmt.Printf("%c %[1]v\n", c)
		word += string(c)
	}
	fmt.Println(word)
}
