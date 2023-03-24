package main

import "fmt"

func main() {
	arrays()
	slices()
}

func arrays() {
	var a [2]string

	a[0] = "Hello"
	a[1] = "World"

	numbers := [5]int{1, 2, 3, 4, 5}

	fmt.Println(a[0], a[1])
	fmt.Println(numbers, numbers[3])
}

func slices() {
	// An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

	// The type []T is a slice with elements of type T.

	// A slice is formed by specifying two indices, a low and high bound, separated by a colon:

	// a[low : high]

	// This selects a half-open range which includes the first element, but excludes the last one.

	// The following expression creates a slice which includes elements 1 through 3 of a:

	// a[1:4]

	numbers := []int{1, 2, 3, 4, 5}
	var s []int = numbers[1:4] // 1 start index, 4 end index (not included)
	fmt.Println(s)

	// Slices are like references to arrays

	// A slice does not store any data, it just describes a section of an underlying array.

	// Changing the elements of a slice modifies the corresponding elements of its underlying array.

	// Other slices that share the same underlying array will see those changes.

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)
	a := names[0:2] // Jhon, Paul
	b := names[1:3] // Paul, Goerge

	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// A slice literal is like an array literal without the length.

	// This is an array literal:

	// [3]bool{true, true, false}

	// And this creates the same array as above, then builds a slice that references it:

	// []bool{true, true, false}

}
