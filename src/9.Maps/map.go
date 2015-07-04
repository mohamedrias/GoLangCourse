package main

import (
	"fmt"
	"sort"	
)

func main() {
	
	/*
		map[KeyType]ValueType
		KeyType can be of any type which is comparable.
		ValueType can be anything including another map itself
		
		The following is a map with keyType string and ValueType int
		Maps are reference types like pointers and slices. It's always pass by reference.
	*/
	var referenceData = map[string]int {
		"admin": 1,
		"moderator": 2,
		"editor": 3,
		"user": 4,
	}
	
	/*
		In the below syntax, map is only declared and not initalized.
		So the value of m will be nil and it doesn't point to any initialized map at all
		
		In case you try to write to a nil map then it will through a runtime panic.
	*/
	
	var m map[string] int;
	
	/*
		Why and How to make map instead of using literal?
		
		If you want to just initialize the map and don't want to assign the variable before hand
		then use make
		
		make function allocates and initializes the data structure and returns a map value
	*/
	
	namesMap = make(map[string] string)
	// namesMap = map[string] string{} another equivalent to make
	
	namesMap["name"] = "Mohamed Rias"
	
	name := namesMap["name"]
	
	/*
		To delete a key from the map
	*/
	
	delete(namesMap, "name")
	
	/*
		To check if the key exists in map
	*/
	
	_, ok := namesMap["name"]
	

	/*
		Iterating maps is similar to that of arrays & slices.
	*/
	for key,value := range referenceData {
		fmt.Println(key,value)
	}
	
	/*
		Skipping the key/value
		_ is used to skip the value. It can be used not only in the map
		but also in terms of functions as well.
	*/
	
	for _,value := range referenceData {
		fmt.Println(value)
	}
	
	/*
		In Map, the iteration order is not guaranteed.
		If the iteration order has to be maintained, then we've to
		manually maintain the order
	*/
	
	

	var employeesMap map[int]string{}
	var keys []string // slice for maintaining the order
	for k := range employeesMap {
	    keys = append(keys, k) // apending all keys to the slice
	}
	sort.Strings(keys) // sort the keys as required
	for _, k := range keys {
	    fmt.Println("Key:", k, "Value:", employeesMap[k])
	}
	
}