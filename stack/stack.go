package stack

import "errors"

// Here, the Stack type is defined to be a slice (i.e., a reference to a variable-length array) of interface{} values.
// This is considered to be different from a plain []interface{}.
type Stack []interface{}

// This is a method definition.
// func (variable type) method_name() return_value { ... }.
// In this example, stack is the RECEIVER.
// In other languages the receiver is typically called 'this' or 'self'.
// It is not considered good Go coding style to use 'this' or 'self' as receiver names.
func (stack Stack) Len() int {
	return len(stack)
}

func (stack Stack) Cap() int {
	return cap(stack)
}

func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}

// The Push method is called on a pointer to a Stack and is passed a value (x) of any type.
func (stack *Stack) Push(x interface{}) {
	// Is append smart enough to determine where to append the value for *stack?

	// Access the value of stack by dereferencing
	*stack = append(*stack, x)
}

// The Top() method returns the top value of the stack
func (stack Stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("can't Top() an empty stack")
	}
	return stack[len(stack)-1], nil
}

// When Go creates a value, it is guaranteed to be its zero value 
// (huge advantage over other programming languages as this eliminates a lot of errors)
// We can optinally write a construction function and call it explicitly

// Pop() method must specify a pointer as the receiver since it modifies the stack by removing the returned item
// 
func (stack *Stack) Pop() (interface{}, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("can't Pop() an empty stack")
	}
	// Grab the top value of the stack (the last item)
	x := theStack[len(theStack)-1]

	// Use the [] slice operator to set the stack to [0:len-1] (all values except last value we just popped from stack)
	// this can be written as [:len-1] or [0:len-1]
	*stack = theStack[:len(theStack)-1]
	return x, nil
}