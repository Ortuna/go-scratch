package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func printPerson(p Person) {
	printPersonPointer(&p)
}

func printPersonPointer(p *Person) {
	fmt.Println(p.firstName)
	fmt.Println(p.lastName)
}

func main() {
	p := new(Person)
	p.firstName = "test"
	p.lastName = "name"
	printPersonPointer(p)

	var v Person
	v.firstName = "test2"
	v.lastName = "name2"

	printPerson(v)
}
