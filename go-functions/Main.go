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

	// s := sum(1, 2, 3, 4, 5)
	// fmt.Println("The sum is", s)

	// d, err := divide(5.0, 3.0)
	// if err != nil { // standard test
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(d)

	// func() {
	// 	msg := "Hello Go!"
	// 	fmt.Println(msg)
	// }()

	// for i := 0; i < 5; i++ { // not good for asynchronous
	// 	func() {
	// 		fmt.Println(i)
	// 	}()
	// }

	// for i := 0; i < 5; i++ { // best practise (asynchronous)
	// 	func(i int) {
	// 		fmt.Println(i)
	// 	}(i)
	// }

	// f := func() { // var f func() {code here}
	// 	fmt.Println("Hello Go!")
	// }
	// f()

	// divide(5.0, 3.0) -> won't work
	// because till this point "divide" function is not declared yet
	// var divide func(float64, float64) (float64, error)
	// divide = func(a, b float64) (float64, error) {
	// 	if b == 0.0 {
	// 		return 0.0, fmt.Errorf("Cannot divide by zero")
	// 	}
	// 	return a / b, nil
	// }
	// d, err := divide(5.0, 3.0)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(d)

	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
	fmt.Println("The new name is:", g.name)
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
//  // Automatically promoted from local memory (stack) to shared memory (heap)
// }

// // syntactic sugar for declaring result variable (initialize with 0)
// // so this variable is gonna be available in the scope of our "sum" function
// // and then that value is gonna be implicitly return
// // but this is actually not done very often in Go language
// // and my suspision is because it can be a little bit confusing to read
// // because your return variables are declared way up there in the top of the function
// // and your actual return is down there on the bottom
// // so if you're reding this code and you're trying to figure it out
// // what this function is actually gonna return
// // you have to come all the way back up to the function signature
// func sum(values ...int) (result int) {
// 	fmt.Println(values)
// 	for _, v := range values {
// 		result += v
// 	}
// 	return
// }

// func divide(a, b float64) float64 {
// 	if b == 0.0 {
// 		panic("Cannot provide zero as second value")
// 	} // but this will stop execution :/
// 	return a / b
// }

// so instead of that
// func divide(a, b float64) (float64, error) { // we can return as many values as we want
// 	if b == 0.0 {
// 		return 0.0, fmt.Errorf("Cannot divide by zero")
// 	}
// 	return a / b, nil
//  // The (result type, error) paradigm is a very common idiom
// 	// this is very idiomatic go
// 	// commonly used error handling
// }

type greeter struct {
	greeting string
	name     string
}

// func (g greeter) greet() { // method
// 	fmt.Println(g.greeting, g.name)
// 	g.name = ""
// }

func (g *greeter) greet() { // method
	fmt.Println(g.greeting, g.name)
	g.name = ""
}
