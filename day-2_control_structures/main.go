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

	age = 35

	switch {
	case age < 18:
		fmt.Println("You are a child")
	case age >= 18 && age < elder_age:
		fmt.Println("You are an adult")
	default:
		fmt.Println("You are an elder")
	}

	switch age {
	case 17:
		fmt.Println("You are 17")
	case 18:
		fmt.Println("You are 18")
	case 35:
		fmt.Println("You are 35")
	default:
		fmt.Println("You are not 17, 18 or 35")
	}
}
