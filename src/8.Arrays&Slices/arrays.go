package main

import "fmt"

func main() {
	// Array notation
	// Need to specify the type of array and the array data
	array := [...]int32 {1,2,3,4}
	for _,x := range array {
		fmt.Println(x)
	}
	
	// Dynamic assignment of values
	// The default value is assigned to other indices in array
	
	array1 := [10]int32{}
	
	array1[0] = 1
	array1[9] = 3

	for _,x := range array1 {
		fmt.Println(x)
	}
	
	stringArray := [5]string {}
	
	stringArray[1] = "Rias"
	stringArray[4] = "Oye"
	for _,x := range stringArray {
		fmt.Println(x)
	}
	
	
}