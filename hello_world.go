package main

import "fmt"

func main() {
	var mapz = map[string]int{
		"test": 123,
		"set":  123,
	}

	fmt.Printf("%d", mapz["test"])
}
