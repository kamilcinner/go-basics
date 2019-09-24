package main

import "fmt"

// const a int16 = 27

// const (
// 	a = iota // this is a counter that can we use when we're creating what we call enumerated constants
// 	b        // = iota
// 	c        // = iota
// )

// const (
// 	a2 = iota
// )

// const ( // we can also write "_ = iota" and then compilers will generate the walue and thrown this away (when we don't need that 0 error handling)
// 	errorSpecialist = iota // to prevent bugs caused by default variable initialization
// 	catSpecialist
// 	dogSpecialist
// 	snakeSpecialist
// )

// const (
// 	_ = iota + 5
// 	catSpecialist
// 	dogSpecialist
// 	snakeSpecialist
// )

// const (
// 	_  = iota             // ignore first value by assigning to blank identifier
// 	KB = 1 << (10 * iota) // shifting by 1 is multiplying by 2
// 	MB
// 	GB
// 	TB
// 	PB
// 	EB
// 	ZB
// 	YB
// )

// coding access roles into a byte
const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	// const myConst int = 42
	// // can't do this (can't set your constant equal to something that has to be determinated in runtime) -> const myConst int = math.Sin(1.57)
	// fmt.Printf("%v, %T\n", myConst, myConst)

	// const a int = 14 // const shadowing
	// const b string = "foo"
	// const c float32 = 3.14
	// const d bool = true
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)

	// const a int = 42
	// var b int = 27
	// fmt.Printf("%v, %T\n", a+b, a+b)

	// this works
	// const a = 42						// (compiler)-> const a = 42
	// var b int16 = 27					// (compiler)-> var b int16 = 27
	// fmt.Printf("%v, %T\n", a+b, a+b)	// (compiler)-> fmt.Printf("%v, %T\n", 42+b, 42+b)

	// fmt.Printf("%v, %T", a, a)
	// fmt.Printf("%v, %T", b, b)
	// fmt.Printf("%v, %T", c, c)
	// fmt.Printf("%v, %T", a2, a2)

	// just enums :)
	// var specialistType int = catSpecialist
	// fmt.Printf("%v\n", specialistType == catSpecialist)

	// fmt.Printf("%v\n", catSpecialist) // 6

	// fileSize := 4000000000.
	// fmt.Printf("%.2fGB\n", fileSize/GB)

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ? %v\n", isHeadquarters&roles == isHeadquarters)
}
