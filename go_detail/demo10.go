package main

import "fmt"

func main() {

}

var a = []int{1, 2, 3}

func copy1() {
	var b = make([]int, len(a))
	copy(b, a)
}

func copy2() {
	var b []int
	b = append(a[:0:0], a...)
	fmt.Print(b)
}
