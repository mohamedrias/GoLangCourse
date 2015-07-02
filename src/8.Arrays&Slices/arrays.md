Arrays in GO have some peculiar characterictics when compared to C arrays

1. Arrays are values and assigning one array to another is always by value and not reference
2. Especially when we pass array to a function and modify the array inside that function, it will not affect the original array as it's by value and not reference
3. The size of array is part of it's type. 
4. The size of array is fixed and need to be mentioned during initialization

