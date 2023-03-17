package utils

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func RandomData(data interface{}, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	fmt.Printf("%v %d\n", data, i)
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}