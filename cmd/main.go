package main

import (
	"fmt"
	"sync"
	"github.com/whauzan/hacktiv8-fga/utils"
)

func main() {
	fmt.Println("Hello! Wahyu here ^-^")

	// Kondisi pertama, data random
	fmt.Printf("\nData Random\n\n")
	coba := []interface{}{"coba1", "coba2", "coba3"}
	bisa := []interface{}{"bisa1", "bisa2", "bisa3"}

	wg := sync.WaitGroup{}

	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go utils.RandomData(coba, &wg, i)
		go utils.RandomData(bisa, &wg, i)
	}

	wg.Wait()

	// Kondisi kedua, data terurut
	fmt.Printf("\nData Terurut\n\n")
	chanCoba := make(chan string)
	chanBisa := make(chan string)

	lockCoba := sync.Mutex{}
	lockBisa := sync.Mutex{}

	for i := 1; i <= 4; i++ {
		go utils.OrderedData(coba, chanCoba, i, &lockCoba)
		go utils.OrderedData(bisa, chanBisa, i, &lockBisa)
	}

	for i := 1; i <= 4; i++ {
		fmt.Printf(<-chanCoba)
		fmt.Printf(<-chanBisa)
	}

	close(chanCoba)
	close(chanBisa)
}
