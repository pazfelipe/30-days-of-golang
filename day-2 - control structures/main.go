package main

import "fmt"

func main() {
	var age int = 17
	var elder_age int = 60

	if age >= elder_age {
		fmt.Println("You are an elder")
	} else if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a child")
	}
}
