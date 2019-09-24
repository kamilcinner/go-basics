package main

import (
	"fmt"
	"strconv"
)

// on the package level we can only use full declaration syntax for variables

// first letter lower-case at the package level meann scooped to the package
// first letter upper-case at the package level means exported and globally visible
// block-scoope menas { only visible between brackets}

var i int = 27

var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

func main() {
	//i = 42
	//k := 99.

	fmt.Println(i)
	var i int = 42 // shadowing - variable in the most inner scope actually takes precedence
	fmt.Printf("%v, %T\n", i, i)

	var j float32
	j = float32(i)
	fmt.Printf("%v, %T\n", j, j)

	var k string
	k = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", k, k)
}
