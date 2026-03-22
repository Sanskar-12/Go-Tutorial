package main

import "fmt"

func main() {
	// Slices

	var arr1 []int // slice in go (slices are dynamically sized arrays)

	arr1=append(arr1, 1,2,3,4)

	fmt.Println("Array 1 ",arr1)

	// Arrays

	var arr2 [4]int // Array are of fixed size

	arr2[0]=1
	arr2[1]=2
	arr2[2]=3
	arr2[3]=4

	fmt.Println("Array 2 ",arr2)
}
