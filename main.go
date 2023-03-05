package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func testPrint(id int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("[%d] counting %d\n", id, i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	var sharedResource string
	var mutex sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mutex.Lock()
			fmt.Println("Previous value: ", sharedResource)
			sharedResource = fmt.Sprintf("key owned by: [%d]", id)
			fmt.Println("Current value: ", sharedResource)
			mutex.Unlock()
			// testPrint(i)
		}(i)
	}

	wg.Wait()
	fmt.Println("final resourc: ", sharedResource)
}
