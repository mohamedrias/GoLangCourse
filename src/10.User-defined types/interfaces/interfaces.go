package main

import (
	"fmt"
)

/**
	Interface
	type interface name followed by interface keyword
*/
type AwesomeDeveloper interface {
	// Interface API methods
	Codes()
}

// Random struct type
type Developer struct {
	name string
}

/*
	Implement the interface.
	Need to use value Developer instead of *Developer for it to work
*/
func (dev Developer) Codes() {
	fmt.Println("Awesome", dev.name)
}

/*
	Generic methods to use across multiple interfaces
*/
func judge(d AwesomeDeveloper) {
	d.Codes()
}

func main() {
	rias := Developer{"Rias"}
	judge(rias)
	rias.Codes()
}
