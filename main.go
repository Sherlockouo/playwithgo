package main

import (
	"fmt"
	"time"
)

func ch(done chan struct{}) {
	fmt.Printf("%v \n", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Printf("%v \n", time.Now())
	done <- struct{}{}
}
func main() {
	done := make(chan struct{})

	go ch(done)

	<-done
	println("shit ")
}
