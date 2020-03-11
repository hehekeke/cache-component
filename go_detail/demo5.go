package main

func main() {
	var n uint = 10
	const N uint = 10
	var x byte = (1 << n) / 100
	var y byte = (1 << N) / 100
	println(x, y)

}
