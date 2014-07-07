package main

import "fmt"

func main() {
	ints := make([]int, 100)
	for i := 0; i < len(ints); i++ {
		fmt.Printf("%d\n", i)
	}
}
