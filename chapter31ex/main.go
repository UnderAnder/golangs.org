package main

import (
	"fmt"
	"strings"
)

func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "goodbye all", "goodbye all", "hello", "hello", "hello", "hello", "goodbye all"} {
		downstream <- v
	}
	close(downstream)
}

func filterCutOnWords(upstream, downstream chan string) {
	for v := range upstream {
		w := strings.Fields(v)
		for _, vv := range w{
			downstream <- vv
		}
	}
	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	prev := ""
	for v := range upstream {
		if v != prev {
			downstream <- v
			prev = v
		}
	}
	close(downstream)
}

func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}

func main() {

	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterCutOnWords(c0, c1)
	printGopher(c1)
}
