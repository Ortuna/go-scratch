package main

import "fmt"

func multiReturn() (returnInt int) {
	returnInt = 50
	return
}

func main() {
	x := multiReturn()
	fmt.Printf("x %d", x)
}
