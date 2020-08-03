package main

import (
	"fmt"
	"strings"
)

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	word := ""

	for i := 0; i < len(cipherText); i++ {
		// 'A' = 0, 'Z' = 25

		c := cipherText[i%len(cipherText)] - (keyword[i%len(keyword)] - 'A')
		if c < 'A' {
			c += 26
		}

		fmt.Printf("%c %[1]v\n", c)
		word += string(c)
	}
	fmt.Println(word)

	message := "your message goes here"
	result := ""
	keyIndex := 0
	message = strings.Replace(message, " ", "", -1)
	message = strings.ToUpper(message)
	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'A' && c <= 'Z' {
			// A=0, B=1, ... Z=25
			c -= 'A'
			k := keyword[keyIndex] - 'A'

			// зашифрованная буква + ключевая буква
			c = (c+k)%26 + 'A'

			// увеличить keyIndex
			keyIndex++
			keyIndex %= len(keyword)
		}
		result += string(c)
	}

	fmt.Println(result)

}
