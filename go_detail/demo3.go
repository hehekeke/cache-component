package main

import "os"

func main() {
	defer println("bye")
	os.Exit(1)
}
