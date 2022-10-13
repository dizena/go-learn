package main

import (
	"fmt"
	"time"
)

func say11(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say11("world")
	say11("hello")
}
