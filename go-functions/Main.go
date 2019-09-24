package main

import "fmt"

func main() {
	// sayMessage("Hello Go!")

	// for i := 0; i < 5; i++ {
	// 	sayMessage("Hello Go!", i)
	// }

	// sayGreetings("Hello", "Stacey")

	// greeting := "Hello"
	// name := "Stacey"
	// sayGreeting(&greeting, &name)
	// fmt.Println(name)

	s := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is", s)
}

// func sayMessage(msg string) {
// 	fmt.Println(msg)
// }

// func sayMessage(msg string, idx int) {
// 	fmt.Println(msg)
// 	fmt.Println("The value of the index is", idx)
// }

// // we're working with the copy
// func sayGreeting(greeting, name string) { // equals (greeting string, name string)
// 	fmt.Println(greeting, name)
// 	name = "Ted"
// 	fmt.Println(name)
// }

// // we're working with the original
// // more efficient because not need to copy
// func sayGreeting(greeting, name *string) { // equals (greeting string, name string)
// 	fmt.Println(*greeting, *name)
// 	*name = "Ted"
// 	fmt.Println(*name)
// }

// func sum(msg string, values ...int) { // variadic parameter has to be the last
// 	fmt.Println(values)
// 	result := 0
// 	for _, v := range values {
// 		result += v
// 	}
// 	fmt.Println(msg, result)
// }

// func sum(values ...int) int {
// 	fmt.Println(values)
// 	result := 0
// 	for _, v := range values {
// 		result += v
// 	}
// 	return result
// }

// func sum(values ...int) *int {
// 	fmt.Println(values)
// 	result := 0
// 	for _, v := range values {
// 		result += v
// 	}
// 	return &result
// 	// Go will automatically promote this value for you
// 	// to be on the shared memory in the computer
// 	// what's also called the heap memory
// }

// syntactic sugar for declaring result variable
// so this variable is gonna be available in the scope of our "sum" function
// and then that value is gonna be implicitly return
// but this is actually not done very often in Go language
// and my suspision is because it can be a little bit confusing to read
// because your return variables are declared way up there in the top of the function
// and your actual return is down there on the bottom
// so if you're reding this code and you're trying to figure it out
// what this function is actually gonna return
// you have to come all the way back up to the function signature
func sum(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}
