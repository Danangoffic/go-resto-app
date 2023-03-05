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

func greet(c chan string) {
	name := <-c // set name as a channel
	fmt.Println("Hello ", name)
}

func greetUntilOut(c chan string, out chan int) {
	for {
		select {
		case name := <-c:
			fmt.Println("Hello ", name)
		case <-out:
			fmt.Println("Ext greeter")
			return
		}
	}
}

func nameReceiver(c chan string, out chan int) {
	for {
		select {
		case name, more := <-c:
			if more {
				fmt.Println("Hello ", name)
			} else {
				fmt.Println("Received all data")
				out <- 0
			}

		}
	}
}

func nameProducer(c chan string) {
	c <- "World" // set channel to "World" string
	c <- "Banana"
	c <- "Apple"
}

func main() {
	c := make(chan string) // channel dengan string
	out := make(chan int)  // channel dengan tipe int
	// go greet(c)
	// go greetUntilOut(c, out)
	go nameReceiver(c, out)
	nameProducer(c)
	close(c)
	// c <- "World" // set channel to "World" string
	// c <- "Banana"
	// c <- "Apple"

	<-out
	// out <- 0 // greeter exit
	// time.Sleep(1 * time.Second)
}

func goroutines() {
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
