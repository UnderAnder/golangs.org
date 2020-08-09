package main

import (
	"fmt"
	"strings"
)

/* func sleepyGopher(id int, c chan int) { // Объявляет канал как аргумент
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	fmt.Println("... ", id, " snore ...")
	c <- id // Отправляет значение обратно к main
} */

func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downstream <- v
	}
	close(downstream)
}

func filterGopher(upstream, downstream chan string) {
	/*     for {
	        item, ok := <-upstream
	        if !ok {
	            close(downstream)
	            return
	        }
	        if !strings.Contains(item, "bad") {
	            downstream <- item
	        }
		} */
	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
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
	go filterGopher(c0, c1)
	printGopher(c1)

	/* 	c := make(chan int) // Делает канал для связи
	   	for i := 0; i < 5; i++ {
	   		go sleepyGopher(i, c)
	   	}
	   	for i := 0; i < 5; i++ {
	   		gopherID := <-c // Получает значение от канала
	   		fmt.Println("gopher ", gopherID, " has finished sleeping")
	   	}

	   	timeout := time.After(2 * time.Second)
	   	for i := 0; i < 5; i++ {
	   		select { // Оператор select
	   		case gopherID := <-c: // Ждет, когда проснется гофер
	   			fmt.Println("gopher ", gopherID, " has finished sleeping")
	   		case <-timeout: // Ждет окончания времени
	   			fmt.Println("my patience ran out")
	   			return // Сдается и возвращается
	   		}
	   	} */

	///////////////////////////////////

}
