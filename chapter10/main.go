package main

import "fmt"

//=============================================
//==================ROT13======================
//=============================================
/*Зашифруйте сообщение на испанском: “Hola Estación Espacial Internacional”  через ROT13.
Модифицируйте Листинг 7 с использованием ключевого слова range.
Теперь, когда вы используете ROT13 c испанским текстом, ударение над буквами сохраняется.
*/
func main() {
	message := "Hola Estación Espacial Internaciona"

	for _, c := range message {
		if c >= 'a' && c <= 'z' {
			c += 13
			if c > 'z' {
				c -= 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c += 13
			if c > 'Z' {
				c -= 26
			}
		}
		fmt.Printf("%c", c)

	}
	/* 	for i := 0; i < len(message); i++ { // Итерирует каждый символ ASCII
		c := message[i]
		if c >= 'a' && c <= 'z' { // Оставляет оригинальную пунктуацию и пробелы
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}*/
}

/*
Расшифруйте цитату Юлия Цезаря: L fdph, L vdz, L frqtxhuhg.

Ваша программа должна будет сдвинуть буквы верхнего и нижнего регистра на -3.
Помните, что ‘a’ становится ‘x’, ‘b’ становится ‘y’, а ‘c’ становится ‘z’.
То же самое происходит с буквами верхнего регистра.
*/
/* func main() {
	message := "L fdph, L vdz, L frqtxhuhg."

	for i := 0; i < len(message); i++ { // Итерирует каждый символ ASCII
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c -= 3
			if c < 'a' {
				c += 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c -= 3
			if c < 'A' {
				c += 26
			}
		}
		fmt.Printf("%c", c)
	}

} */
