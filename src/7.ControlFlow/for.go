package main

import "fmt"

func main() {
	/*
		for loop in Go unifies for, while and do while.
		There are three forms of for loop and it helps to cover
		all the variants of loop
	*/

	/**
	Equivalent to for(;;) in C
	While (true){} infinite loop
	*/
	for {
		fmt.Println("Hello World")
		// Breaking it just for brevity and preventing the infinite loop
		break;
	}

	/**
	Typical for loop syntax and behavior
	*/
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}

	
	/*
		If you only need the second item in the range (the value),
		use the blank identifier, an underscore, to discard the first:
		
		range is used to create a range from String, Array, Map, slice
	*/
	
	name := "MohamedRias"
	for _,letter := range name {
		fmt.Println(letter, string(letter))
	}
}
