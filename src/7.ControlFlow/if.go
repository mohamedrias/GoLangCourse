package main

import (
	"fmt"
)

func main() {
	//If statements are same as C language but doesn't need to include () braces
	x := "hi there"
	if x == "hi there" {
		fmt.Println("Hi There :P")
	}

	i := 20
	/*
		If statements can take optional initialization parameters as well.
		This is different compared to the traditional C approach
	*/
	if y := 10; i < y {
		fmt.Printf("%d is less than %d\n", i, y)
	} else {
		fmt.Printf("%d is greater than %d\n", i, y)
	}

	/**
	Can initialize multiple arguments as well.
	It keeps the if condition very simple
	*/

	if x, y := 21, 12; x%3 == 0 && y%7 == 0 {
		fmt.Printf("X / y is not a factor")
	} else if x%7 == 0 && y%3 == 0 {
		fmt.Printf("Both numbers are modulus of 3 & 7")
	}
}
