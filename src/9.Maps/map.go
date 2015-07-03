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
}