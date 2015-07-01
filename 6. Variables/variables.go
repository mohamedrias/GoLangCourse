package main

import (
	"fmt"
)

func main() {
	
	/*
	var keyword represents variable
	a,b represents variable names
	bool is the data type
	
	It will store only TRUE/FALSE
	*/
	var a, b bool
	a = true
	b = false
	
	fmt.Println(a,b)
	
	/**
	float32 and float64 are used to allocate appropriate memory to the 
	decimal numbers
	*/
	
	var c float32
	c = 3.14
	fmt.Println(c)

	/**
	Different type of integers are available. By default it's 32 bit
	If required we can specify it to be 64 bits as well
	*/	
	var i1 int
	var	i2 int64 
	var	i3 uint32
		
	fmt.Println(i1, i2, i3)
		
	/**
	TODO: refer more about rune
	
	rune is equivalent to int32 but supports unicode characters
	Inside printf statements refer it as %#U 	
	*/
	//var r rune
	//r = "⌘"
	//fmt.Println(r)
	
	
	b = 5 < 10
	fmt.Println(5)
	
	var f2 float32
	f2 = 3*(22/7)
	fmt.Println(f2)
	
	
	var s string
	s = "Hello World"
	fmt.Println(s)
	
	//To check what is there in position 4
	// refer strings package which returns a lot of useful methods
}