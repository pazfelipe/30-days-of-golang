package main

import "fmt"

func main() {
	var timezone = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
	}

	fmt.Println(timezone)
	fmt.Println(timezone["EST"])

	timezone["CET"] = 1 * 60 * 60

	fmt.Println(timezone)

	attend := map[string]bool{
		"Bob":  true,
		"Mary": true,
	}

	fmt.Println(attend)

	delete(attend, "Bob")

	fmt.Println(attend)

	// If key is in attend, ok is true. If not, ok is false.

	// If key is not in the map, then elem is the zero value for the map's element type.
	elem, ok := attend["Mary"]
	fmt.Println(elem, ok)

	elem, ok = attend["Bob"]
	fmt.Println(elem, ok)
}
