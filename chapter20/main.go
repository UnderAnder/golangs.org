package main

import (
	"fmt"
	"strings"
)

func listOfWords(text string) []string {
	text = strings.ToLower(text)
	text = strings.Trim(text, ".,!;")
	return strings.Fields(text)
}

func main() {
	text := "As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man."

	frequency := make(map[string]int)

	for _, t := range listOfWords(text) { // Итерирует через срез (индекс, значение)
		frequency[t]++
	}

	for t, num := range frequency { // Итерирует через карту (ключ, значение)
		fmt.Printf("%v встречается %d раз(а) \n", t, num)
	}

}
