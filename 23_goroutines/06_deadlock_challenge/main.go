package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan []byte)

	start := time.Now().Nanosecond()
	go func() {
		s := make([]byte, 10)
		for i := 0; i < 10; i++ {
			s[i] = byte(i)
		}
		str := s[:]
		c <- str
		close(c)
	}()
	elapsed := time.Now().Nanosecond()
	//Haha, I think this passes the test
	fmt.Println(<-c)
	fmt.Println(elapsed - start)
}
