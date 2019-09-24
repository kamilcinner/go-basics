package main

import "fmt"

func main() {
	// var n bool = true
	// fmt.Printf("%v, %T\n", n, n)

	// k := 1 == 1
	// l := 1 == 2
	// fmt.Printf("%v, %T\n", k, k)
	// fmt.Printf("%v, %T\n", l, l)

	// var m bool
	// fmt.Printf("%v, %T\n", m, m)

	// int8, int16, int32, int64

	// uint8, uint16, uint32
	// var n uint16
	// fmt.Printf("%v, %T\n", n, n)

	// a := 10
	// b := 3
	// fmt.Println(a + b)
	// fmt.Println(a - b)
	// fmt.Println(a * b)
	// fmt.Println(a / b)
	// fmt.Println(a % b)

	// we can't add two integers of different types
	// var a int = 10
	// var b int8 = 3
	// fmt.Println(a + int(b))

	// bit operator
	// a := 10             // 1010
	// b := 3              // 0011
	// fmt.Println(a & b)  // 0010
	// fmt.Println(a | b)  // 1011
	// fmt.Println(a ^ b)  // 1001
	// fmt.Println(a &^ b) // 0100

	// bit shifting
	// a := 8              // 2^3
	// fmt.Println(a << 3) // 2^3 * 2^3 = 2^6 = 64
	// fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0 = 1

	// default := will be float64
	// n := 3.14
	// n = 13.7e72
	// n = 2.1E14
	// fmt.Printf("%v, %T\n", n, n)

	// floating point arithmetic operations
	// a := 10.2
	// b := 3.7
	// fmt.Println(a + b)
	// fmt.Println(a - b)
	// fmt.Println(a * b)
	// fmt.Println(a / b)

	// complex number <3
	// var n complex128 = 1 + 2i
	// fmt.Printf("%v, %T\n", n, n)
	// fmt.Printf("%v, %T\n", real(n), real(n))
	// fmt.Printf("%v, %T\n\n", imag(n), imag(n))

	// var m complex128 = complex(5, 12)
	// fmt.Printf("%v, %T\n\n", m, m)

	// a := 1 + 2i
	// b := 2 + 5.2i
	// fmt.Println(a + b)
	// fmt.Println(a - b)
	// fmt.Println(a * b)
	// fmt.Println(a / b)

	// s := "this is a string"
	// fmt.Printf("%v, %T\n", string(s[2]), s[2])

	// s2 := "this is also a string"
	// fmt.Printf("%v, %T\n", s+s2, s+s2)

	// array of uint8s
	// b := []byte(s)
	// fmt.Printf("%v, %T\n", b, b)

	var r rune = 'a'
	fmt.Printf("%v, %T\n", r, r)

}
