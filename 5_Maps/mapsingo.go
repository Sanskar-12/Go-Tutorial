package main

import "fmt"

func main() {
	// make a map
	productPrice := map[string]int {
		"Iphone 11":1000,
		"Iphone 12":2000,
		"Iphone 13":3000,
	}

	fmt.Println("Product price ",productPrice)

	// make a empty map
	customMap := map[string]string {}

	fmt.Println("Custom map ",customMap)

	emptyMap := make(map[string]int)

	// add the key value pair in the map
	emptyMap["key1"]=100
	emptyMap["key2"]=200
	emptyMap["key3"]=300

	fmt.Println("Empty map ",emptyMap)

	for key, value := range emptyMap {
		fmt.Println("Key: ",key," Value: ",value)
	}
}
