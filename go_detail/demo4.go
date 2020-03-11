package main

import "runtime"

func main() {
	c := make(chan int)
	go func() {
		defer close(c)
		defer println("bye")
		runtime.Goexit()
	}()
	<-c
}
