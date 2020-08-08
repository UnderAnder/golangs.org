package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	p, err := url.Parse("https://a b.com/")
	if err != nil {
		if e, ok := err.(*url.Error); ok {
			fmt.Println("Op:", e.Op)   // Выводит: Op: parse
			fmt.Println("URL:", e.URL) // Выводит: URL: https://a b.com/
			fmt.Println("Err:", e.Err) // Выводит: Err: invalid character “ ” in host name
		}
		os.Exit(1)
	}
	fmt.Printf("%#v", p)
}
