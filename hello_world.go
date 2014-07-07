package main

import "fmt"

func multiReturn() (int, string) {
	return 1, "Some string"
}

func main() {
	x, y := multiReturn()
	fmt.Printf("x %d, y %s", x, y)
}
