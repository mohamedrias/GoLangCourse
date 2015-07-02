package main

import "fmt"

func main() {

	var language string
	language = "PHP"

	/**
	No need to insert semicolon as in C
	Mention the variable in switch and write appropriate conditions in cases

	UNlike java/C, case can take conditions instead of simple strings

	No need to specify break. statement after each case
	*/
	switch language {
	case "go":
		fmt.Println("Go is awesome")
	case "Java":
		fmt.Println("Hmm, it's for enterprises")
	case "Node":
		fmt.Println("Event driven")
	case "PHP":
		fmt.Println("Did you chose PHP?")
		// Use fallthrough to continue to the next statement
		fallthrough
	default:
		fmt.Println("That's defacto")
	}

	/*
		No need to specify an expression in the Switch statement.
		If no expresion is provided, by default it assumes it to be true
		So ideally any if else if condition can be put in a switch statement easily


	*/
	c := 7
	switch {
	case 8 <= c && c <= 9:
		fmt.Println(c)
	case 89 <= c && c <= 56:
		// Condition
	case 76 <= c && c <= 78:
		// Condition
		//fallthrough
	default:
	}

	// How to use multiple || conditions on the same case

	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		fmt.Println("Contains any one of the characters")
	default:
		fmt.Println("Not matching")
	}

	//TODO: Create example for Type Switch
}
