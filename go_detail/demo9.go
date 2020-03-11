package main

var (
	_ = f3("w", x1)

	x1 = f3("x", z)

	y = f3("y", x1)

	z = f3("z")
)

func f3(s string, deps ...int) int {
	println(s)
	return 0
}

func main() {

}
