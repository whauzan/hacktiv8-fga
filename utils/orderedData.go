package utils

import (
	"fmt"
	"sync"
)

func OrderedData(data interface{}, c chan string, i int, mutex *sync.Mutex) {
	mutex.Lock()
	c <- fmt.Sprintf("%v %d\n", data, i)
	mutex.Unlock()
}