package main

func f1() bool {
	return false
}

func main() {
	switch f1() 
	{
	case true:println(true)
	case false:println(false)
	}
}
