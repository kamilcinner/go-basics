package main

import (
	"fmt"
)

func main() {
	// if true {
	// 	fmt.Println("The test is true")
	// }

	// statePopulations := map[string]int{
	// 	"California":   39250017,
	// 	"Texas":        27862596,
	// 	"Florida":      20612439,
	// 	"New York":     19745289,
	// 	"Pennsylvania": 12802503,
	// 	"Illinois":     12801539,
	// 	"Ohio":         11614373,
	// }
	// if pop, ok := statePopulations["Florida"]; ok { // pop is defined only for if scope
	// 	fmt.Println(pop)
	// }

	// number := 50
	// guess := 30
	// if guess < 1 || returnTrue() || guess > 100 { // if first left statement returns true then GO doen't have to execute any more code
	// 	fmt.Println("The guess must be between 1 and 100!")
	// }
	// if guess >= 1 && guess <= 100 { // same rule here
	// 	if guess < number {
	// 		fmt.Println("Too low")
	// 	}
	// 	if guess > number {
	// 		fmt.Println("too high")
	// 	}
	// 	if guess == number {
	// 		fmt.Println("You got it!")
	// 	}
	// 	fmt.Println(number <= guess, number >= guess, number != guess)
	// }
	// fmt.Println(!true)

	// number := 50
	// guess := 30
	// if guess < 1 || guess > 100 {
	// 	fmt.Println("The guess must be between 1 and 100!")
	// } else {
	// 	if guess < number {
	// 		fmt.Println("Too low")
	// 	}
	// 	if guess > number {
	// 		fmt.Println("too high")
	// 	}
	// 	if guess == number {
	// 		fmt.Println("You got it!")
	// 	}
	// 	fmt.Println(number <= guess, number >= guess, number != guess)
	// }

	// number := 50
	// guess := -300
	// if guess < 1 {
	// 	fmt.Println("The guess must be greater than 1!")
	// } else if guess > 100 {
	// 	fmt.Println("The guess must be less than 100!")
	// } else {
	// 	if guess < number {
	// 		fmt.Println("Too low")
	// 	}
	// 	if guess > number {
	// 		fmt.Println("too high")
	// 	}
	// 	if guess == number {
	// 		fmt.Println("You got it!")
	// 	}
	// 	fmt.Println(number <= guess, number >= guess, number != guess)
	// }

	// // not a great idea
	// myNum := 0.123
	// if myNum == math.Pow(math.Sqrt(myNum), 2) {
	// 	fmt.Println("These are the same")
	// } else {
	// 	fmt.Println("These are different")
	// }

	// // better way with error parameter
	// myNum := 0.123
	// if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
	// 	fmt.Println("These are the same")
	// } else {
	// 	fmt.Println("These are different")
	// }

	// switch 2 {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// default:
	// 	fmt.Println("not one or two")
	// }

	// // can't have overlapping cases
	// switch 2 {
	// case 1, 5, 10:
	// 	fmt.Println("one, five or ten")
	// case 2, 4, 6:
	// 	fmt.Println("two, four or six")
	// default:
	// 	fmt.Println("another number")
	// }

	// switch i := 2 + 3; i {
	// case 1, 5, 10:
	// 	fmt.Println("one, five or ten")
	// case 2, 4, 6:
	// 	fmt.Println("two, four or six")
	// default:
	// 	fmt.Println("another number")
	// }

	// are allowed to overlap
	// first state that evaluetes to true is gona excecute
	// GO has implicite break between our cases
	// instead of implicit fallthrough
	// but we can use word "fallthrough" to change that
	// but also fallthrough is LOGIC LESS !!!
	// i := 10
	// switch {
	// case i <= 10:
	// 	fmt.Println("less than or equal to ten")
	// case i <= 20:
	// 	fmt.Println("less than or equal to twenty")
	// default:
	// 	fmt.Println("greater than twenty")
	// }

	// i := 10
	// switch {
	// case i <= 10:
	// 	fmt.Println("less than or equal to ten")
	// 	fallthrough // this is logic less
	// case i >= 20:
	// 	fmt.Println("less than or equal to twenty")
	// default:
	// 	fmt.Println("greater than twenty")
	// }

	var i interface{} = [3]int{} // interface can take any type of data that we have in GO
	// this won't work -> var i int = 1 (it must be an interface)
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
		break // useful for debugging
		// fmt.Println("This won't print")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is a string")
	case [3]int:
		fmt.Println("i is [3]int")
	default:
		fmt.Println("i is another type")
	}
}

// func returnTrue() bool {
// 	fmt.Println("returning true")
// 	return true
// }
