package main

import "fmt"

type TestMap struct {
	Map map[int]int
}

func main() {
	var testMap = TestMap{
		Map: map[int]int{
			1: 1,
			2: 2,
		},
	}
	var res = make(map[int]int)
	for k, v := range testMap.Map {
		res[k] = v
	}
	//var res = testMap.Map
	fmt.Println(res)

	fmt.Println(testMap.Map)
	res[3] = 3
	fmt.Println(testMap.Map)
	testMap.Map[4] = 4

	fmt.Println(res)
}
