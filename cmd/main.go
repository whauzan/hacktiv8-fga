package main

import (
	"fmt"
)

func main() {
	fmt.Printf("\n\nHello! Wahyu here ^-^\n\n")

	text := "selamat malam"
	counter := make(map[string]int)
	for _, c := range text {
		fmt.Printf("%v\n", string(c))
		counter[string(c)]++
	}

	fmt.Println(counter)
}
