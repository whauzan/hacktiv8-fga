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

    chars := make([]string, 0, len(counter))
    for chr := range counter {
        chars = append(chars, chr)
    }

	fmt.Println(counter)
}