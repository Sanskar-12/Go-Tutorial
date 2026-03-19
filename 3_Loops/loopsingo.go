package main

import "fmt"

func main() {

	// 1st way to use for loops
	for i:=1;i<=10;i++ {
		fmt.Println(i)
	}

	// 2nd way to use for loops using range keyword
	for i := range 3 {
		fmt.Println(i)
	}

	for i, char := range "Sanskar" {
		fmt.Println(i,char)
	}

	// while loop using for loop in GO
	j:=10

	for j>0 {
		// if, else if and else conditional in GO
		if j==3 {
			j--
			continue // will skip that iteration
		} else if j==5 {
			j--
			continue
		} else {
			j--
			continue
		}
		fmt.Println(j)
		j--
	}

	// infinite loop and break keyword to break the loop
	for {
		fmt.Println("Infinite loop")
		break // will break the loop
	}
}