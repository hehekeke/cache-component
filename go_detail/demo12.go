package main

import "fmt"

func main() {
	fmt.Printf("%[2]v%[1]v%[2]v%[1]v", "o", "c")
}
