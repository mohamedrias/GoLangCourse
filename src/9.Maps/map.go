package main

import "fmt"

func main() {
	
	var referenceData = map[string]int {
		"admin": 1,
		"moderator": 2,
		"editor": 3,
		"user": 4,
	}

	/*
		Iterating maps is similar to that of arrays & slices
		we can use _ for skipping the key if required
	*/
	for key,value := range referenceData {
		fmt.Println(key,value)
	}
	
	/*
		Skipping the key/value
		_ is used to skip the value. It can be used not only in the map
		but also in terms of functions as well.
	*/
	
	for _,value := range referenceData {
		fmt.Println(value)
	}
}