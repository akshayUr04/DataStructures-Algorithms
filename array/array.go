package main

import "fmt"

func main() {
	arr := make([]int, 2, 3)
	fmt.Println(arr)
	fmt.Println("cap", cap(arr))
	fmt.Println("size", len(arr))
	arr = append(arr, 10)
	arr = append(arr, 11)
	arr = append(arr, 12)
	arr = append(arr, 13)
	arr = append(arr, 14)
	fmt.Println("cap1", cap(arr))
	fmt.Println("size1", len(arr))

	fmt.Println(arr)
}
