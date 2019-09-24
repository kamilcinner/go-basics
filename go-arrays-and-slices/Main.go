package main

import "fmt"

func main() {
	// Arrays are fixed size

	// var students [3]string
	// grades := [...]int{97, 85, 93}
	// fmt.Printf("Grades: %v\n\n", grades)

	// students[0] = "Lisa"
	// students[1] = "Ahmed"
	// students[2] = "Arnold"
	// fmt.Printf("Students: %v\n\n", students)

	// fmt.Printf("Student #1: %v\n\n", students[1])

	// fmt.Printf("Number of Students: %v\n\n", len(students))

	// var identityMatrix [3][3]int
	// identityMatrix[0] = [3]int{1, 0, 0}
	// identityMatrix[1] = [3]int{0, 1, 0}
	// identityMatrix[2] = [3]int{0, 0, 1}
	// fmt.Println(identityMatrix)

	// fmt.Println()

	// a := [...]int{1, 2, 3}
	// b := a // copy
	// b[1] = 5
	// fmt.Println(a)
	// fmt.Println(b)

	// fmt.Println()

	// c := [...]int{1, 2, 3}
	// d := &c // original
	// d[1] = 5
	// fmt.Println(c)
	// fmt.Println(d)

	// Slices

	// a := []int{1, 2, 3}
	// b := a // original
	// b[1] = 5
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n", cap(a))

	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// b := a[:]   // slice of all elements
	// c := a[3:]  // slice from 4th element to end
	// d := a[:6]  // slice first 6 elements
	// e := a[3:6] // slice the 4th, 5th and 6th elements (first is inclusive, last is exclusive)
	// a[5] = 42   // all of them change their value
	// // we can use slices operations also with arrays
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)

	// a := make([]int, 3)
	// // first  arg - type of object that we want to create
	// // second arg - the length of the slice
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n\n", cap(a))

	// b := make([]int, 3, 100)
	// // third arg - the capacity of the slice
	// fmt.Println(b)
	// fmt.Printf("Length: %v\n", len(b))
	// fmt.Printf("Capacity: %v\n", cap(b))

	// a := []int{}
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n\n", cap(a))
	// // copies all existing elements to a new array that has larger size
	// // arrays with larger size will be very expensive in this way
	// // it is good to make a bigger capacity if we know it will be needed
	// a = append(a, 1)
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n\n", cap(a))
	// // capacity will automatic increase if needed
	// a = append(a, 2, 3, 4, 5)
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n\n", cap(a))
	// // someone says that it double the size of the previous array
	// // but this doesn't look so
	// // but be aware of that with powers of two you can actually went with
	// // a lot of memory consumed that you never gona be using
	// // decent first estimate will be beneficial to you

	// // concatenate two slices with spread operator (...)
	// a = append(a, []int{6, 7, 8, 9, 10}...) // ... is a spread operator
	// fmt.Println(a)
	// fmt.Printf("Length: %v\n", len(a))
	// fmt.Printf("Capacity: %v\n\n", cap(a))

	a := []int{1, 2, 3, 4, 5}
	// b := a[:len(a)-1]
	fmt.Printf("a: %v\n", a)
	b := append(a[:2], a[3:]...)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("a: %v\n", a)
	// a has changed because of the reference
}
