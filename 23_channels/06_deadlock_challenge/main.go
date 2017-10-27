package main

import "fmt"

func main() {
	c := make(chan []byte)

	go func() {
		s:=make([]byte, 10)
		for i := 0; i < 10; i++ {
			s[i] = byte(i)
		}
		str := s[:]
		c <- str
		close(c)
	}()
		//Haha, I think this passes the test
		fmt.Println(<-c)
}