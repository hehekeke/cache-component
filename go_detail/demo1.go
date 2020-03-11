package main

func f() bool {
	return false
}

func main() {
	switch f() {
	case true:
		println(true)
	case false:
		println(false)
	}
}
