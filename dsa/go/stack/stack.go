package stack

import "fmt"

var MAX_SIZE = 5
var stacks [5]int
var top = -1

func Run() {
out:
	for {
		fmt.Println("\nImplementation of stack.\n0. Exit\n1. Push\n2. Pop\n3. Display")
		var input int
		fmt.Print("Your Input:")
		fmt.Scanln(&input)
		if input > 3 || input < 0 {
			fmt.Println("Please input 0, 1, 2 or 3")
		}
		switch input {
		case 1:
			var num int
			fmt.Print("Input Number to Push: ")
			fmt.Scanln(&num)
			push(num)
		case 2:
			pop()
		case 3:
			display()
		case 0:
			break out
		}
	}
}
func isFull() bool {
	return top == MAX_SIZE-1
}
func push(num int) {
	if !isFull() {
		top += 1
		stacks[top] = num
		fmt.Println("Pushed ", num, " to the stack.")
	} else {
		fmt.Println("Stack Full.")
	}
}
func isEmpty() bool {
	return top == -1
}
func pop() {
	if !isEmpty() {
		removed := stacks[top]
		stacks[top] = 0
		top -= 1
		fmt.Println("Removed: ", removed)
	} else {
		fmt.Println("Stack already empty.")
	}
}
func display() {
	fmt.Println("\n------- STACK -------")
	for i := len(stacks) - 1; i >= 0; i-- {
		fmt.Println(i+1, ". ", stacks[i])
	}
	fmt.Print("-----------------------")
}

// import (
//     "fmt"
// )
//
// type stack[T any] struct {
//     Push func(T)
//     Pop func() T
//     Length func() int
// }
//
// func Stack[T any]() stack[T] {
//     slice := make([]T, 0)
//     return stack[T]{
//         Push: func(i T) {
//             slice = append(slice, i)
//         },
//         Pop: func() T {
//             res := slice[len(slice)-1]
//             slice = slice[:len(slice)-1]
//             return res
//         },
//         Length: func() int {
//             return len(slice)
//         },
//     }
// }
//
//
// func main() {
//     stack := Stack[string]()
//     stack.Push("this")
//     fmt.Println(stack.Length())
//     fmt.Println(stack.Pop())
// }
