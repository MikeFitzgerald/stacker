/*
	
1. Go supports the creation of custom types (is-a relationships)

2. Go makes aggregation (has-a) very easy

3. Go uses Duck Typing (if it _s like a duck, it's a duck)

4. Go's custom types are based on the fundamental types or structs

5. Go supports both named and unnamed custom types

6. Unnamed custom types with the same structure can be used interchanably but CANNOT have any methods.

7. Named custom types CAN have methods; these methods compose the type's interface.

8. The empty interface is specified as interface{} - think of it kinda like a pointer to any type.

9. Function and method parameters can be of any built-in or custom type
   (example: "pass a value that can read data" instead of "pass this specific type of value")

*/

package main

import (
   "fmt"
   "github.com/MikeFitzgerald/stacker/stack"
)

func main() {
	var haystack stack.Stack
	fmt.Println("Stack empty?", haystack.IsEmpty(), "\n")

	fmt.Println("Adding values...")
	haystack.Push("hay")
	haystack.Push(-15)
	haystack.Push([]string{"pin", "clip", "needle"})
	haystack.Push(81.52)

	fmt.Println("Stack empty?", haystack.IsEmpty(), "\n")

	for {
		item, err := haystack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}

	fmt.Println("\nStack empty?", haystack.IsEmpty(), "\n")	
}