package main

import s "fmt"

/**
*	Main program which will be invoked automatically while running the file
*/
func main() {
	// Variable declaration
	// var keyword, input is variable name and string is the datatype
	var input string;
	
	//Println is used to display message in the console
	s.Println("What's your name?");
	
	//Scanf is used to get the user input.
	s.Scanf("%s", &input);
	
	// String concatenation using , and +
	s.Println("Hello,", input+ "!");
}