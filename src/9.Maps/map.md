1. Maps is a built-in data structure that associate values of one type (the key) with values of another type (the element or value) 

2. The key can be of any type for which the equality operator is defined, such as integers, floating point and complex numbers, strings, pointers, interfaces (as long as the dynamic type supports equality), structs and arrays. 

3. Slices cannot be used as map keys, because equality is not defined on them. 

4. Like slices, maps hold references to an underlying data structure. 

5. If you pass a map to a function that changes the contents of the map, the changes will be visible in the caller.