package main

import _ "fmt"

func main() {
	
	type awesomeDeveloper interface {
		eats()
		codes()
		sleeps()
		repeats()
	}
	
	type dev struct {
		name string
	}
	
	func (d *dev) eats() string {
		fmt.Println("Rias")
		return "Rias"
	}
	
	//TODO: Need to refactor. It's giving error :(
}