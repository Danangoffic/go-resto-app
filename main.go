package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type myService struct {
	cache *cache
}

func (ms *myService) expensiveFunction(key string) string {
	for i := 0; i < 50000; i++ {
		fmt.Printf("Data index ke : %v", i)
	}
	return fmt.Sprint("data-", key, ":success")
}

func (ms *myService) GetData(key string) string {
	if ms.cache != nil {
		if cacheData := ms.cache.get(key); cacheData != "" {
			return cacheData
		}
	}
	result := ms.expensiveFunction(key)
	ms.cache.set(key, result)
	return result
}

type cache struct {
	storage map[string]string
}

func (c *cache) set(key, value string) {
	c.storage[key] = value
}

func (c *cache) get(key string) string {
	v, ok := c.storage[key]
	if !ok {
		return ""
	}
	return v
}

func main() {
	cacher := &cache{
		storage: map[string]string{},
	}
	service := &myService{
		cache: cacher,
	}

	key := "mydata1"

	start := time.Now()
	fmt.Println("calling expensive function")
	result := service.GetData(key)
	fmt.Println("expensive function called, duration: ", time.Since(start))
	fmt.Println(result)

	start = time.Now()
	fmt.Println("calling expensive function")
	result = service.GetData(key)
	fmt.Println("cache expensive function called, duration: ", time.Since(start))
	fmt.Println(result)
}

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

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var total V
	for _, v := range m {
		total += v
	}
	return total
}

type Number interface {
	int | int32 | int64 | float32 | float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var total V
	for _, v := range m {
		total += v
	}
	return total
}

func genericsFunc() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 32.3,
	}

	fmt.Printf("Generic Sum total int and float: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats),
	)
	fmt.Printf("Generic Sum total int and float with type interface: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats),
	)
}

func channels() {
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
