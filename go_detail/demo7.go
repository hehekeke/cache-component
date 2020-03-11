package main

var x, i = []int{1, 2}, 0

func f2() int {
	i = 1
	return 9
}

func main() {
	x[i] = f2()
	println(x[0], x[1])
}
