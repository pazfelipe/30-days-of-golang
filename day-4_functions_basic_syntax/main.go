package main

import "fmt"

func printing() {
	fmt.Println("Hello World")
}

func sum(a int, b int) int {
	return a + b
}

func voidReturn() {
	result := sum(2, 5)
	fmt.Println(result)
}

func arrayParameter(arr []byte, arr2 []int, arr3 [2]string) {
	fmt.Println(arr, len(arr))
	fmt.Println(arr2, len(arr2))
	fmt.Println(arr3, len(arr3))
}

func main() {
	printing()
	voidReturn()
	arrayParameter([]byte{1, 2, 3, 4, 5}, []int{1, 2, 3, 4}, [2]string{"a", "b"})
}
