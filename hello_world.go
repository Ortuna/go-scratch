package main

import "fmt"

func printMessage(flag int) {
	if flag == 0 {
		fmt.Println("printing 0")
	} else if flag == 1 {
		fmt.Println("printing 1")
	}
}

func main() {
	defer printMessage(1)
	fmt.Println("In main")
	return
}
