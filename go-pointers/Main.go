package main

import "fmt"

func main() {
	// var a int = 42
	// var b *int = &a
	// fmt.Println(&a, b)
	// a = 27
	// fmt.Println(a, *b)
	// *b = 14
	// fmt.Println(a, *b)

	// a := [3]int{1, 2, 3}
	// b := &a[0]
	// c := &a[1]
	// fmt.Printf("%v %p %p\n", a, b, c)
	// unsafe package allow pointer arithmetic
	// by default Go it's not allowed to safe language simplicity

	// var ms *myStruct
	// ms = &myStruct{foo: 42}
	// fmt.Println(ms)

	// var ms *myStruct
	// fmt.Println(ms)
	// ms = new(myStruct)
	// fmt.Println(ms)

	// var ms *myStruct
	// ms = new(myStruct)
	// (*ms).foo = 42 // * operator has a lower presidence than the . operator
	// // so we  need () in order to make sure that we dereferencing the ms variable instead of dereferencing ms.foo
	// fmt.Println((*ms).foo)

	// // but with all of these compilers can help us a little bit
	// var ms *myStruct
	// ms = new(myStruct)
	// ms.foo = 42
	// fmt.Println(ms.foo)
	// // same result as upper

	// a := [3]int{1, 2, 3} // array
	// b := a
	// fmt.Println(a, b)
	// a[1] = 42
	// fmt.Println(a, b)

	// a := []int{1, 2, 3} // slice
	// b := a
	// fmt.Println(a, b)
	// a[1] = 42
	// fmt.Println(a, b)

	a := map[string]string{"foo": "bar", "baz": "buz"}
	b := a
	fmt.Println(a, b)
	a["foo"] = "qux"
	fmt.Println(a, b)
}

type myStruct struct {
	foo int
}
