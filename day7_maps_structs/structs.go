package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{
		Name: "Bob",
		Age:  20,
	}

	fmt.Println(person)
	fmt.Println(person.Name)

	fmt.Println("declaring directly", Person{"Jhon", 30})
}
