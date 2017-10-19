package main

import (
	"fmt"
	"sync"
	"time"
	"math/rand"
)
var wg sync.WaitGroup
var count int
var mutex sync.Mutex

func main() {
	//Add four positions in the waitgroup
	wg.Add(4)
	go incrementor("Foo:")
	go incrementor("Bar:")
	go incrementor("Wid:")
	go incrementor("Get:")
	wg.Wait()
	//Print final count after all functions have had an oppourtunity to run
	fmt.Println("Final Counter:", count)
}

func incrementor(s string) {
	for i := 0; i < 100; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		mutex.Lock()
		count++
		fmt.Println(s, i, "Counter:", count)
		mutex.Unlock()
	}
	wg.Done()
}

//go run -race main.go
//vs
//go run main.go