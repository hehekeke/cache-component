package main

var s string

var x2 = []byte{1023: 'x'}
var y1 = []byte{1023: 'y'}

func main() {

}
func fc() {
	s = (" " + string(x1) + string(y1))[1:]

}

func fd() {
	s = string(x2) + string(y1)
}
