package main

import "fmt"

func main() {
	message := "uv vagreangvbany fcnpr fgngvba"

	for i := 0; i < len(message); i++ { // Итерирует каждый символ ASCII
		c := message[i]
		if c >= 'a' && c <= 'z' { // Оставляет оригинальную пунктуацию и пробелы
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
}
