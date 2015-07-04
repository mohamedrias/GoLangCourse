// Struct - user defined data types

package main

import "fmt"

func main() {
	
	/* 
		struct declaration
		Use capital case to identify structs easily.
	*/
	type Employee struct {
		id int
		name string
		company string
	}
	
	/*
		For using employee struct, use pointer notation
	*/
	
	rias := Employee{
		name: "Rias",
		id: 4,
		company: "TCS",	
	}
	
	/*
		Printing struct object
	*/
	
	fmt.Println(rias)
	
	/*
		To Copy one struct into another	
	*/
	
	riasNew := &rias
	
	fmt.Println(riasNew)
	
	riasNew.name = "Mohamed Rias"
	
	fmt.Println(riasNew, rias)
	/*
		Other ways of refering to the struct object
	*/
	
	fmt.Println(Employee{2,"Rafi", "3i Infotech"})
	
	
	fmt.Println(Employee{id: 3, name: "Hiren", company: "Babajob"})
	
	
	/*
		Traditional struct declaration
	*/
	
	alex := Employee{}
	alex.id = 10
	alex.name = "Alex Rattray"
	alex.company = "Babajob"
	
	fmt.Println(alex)
	
}