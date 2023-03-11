package main
import "fmt"

// Use var to declare a variable followed by type and value
var a bool = true
var b int = 10
var c float32 = 10.5

func main() {
	
	var d string = "Hello World"
	
	// User := to declare a variable and assign a value
	// Type is inferred from the value

	e := 10 // int
	f := 10.5 // float64
	g := "Hello World" // string


	// Difference between var and :=
	// var can be used outside a function
	// := can only be used inside a function

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}