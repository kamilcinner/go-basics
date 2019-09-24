package main

import (
	"fmt"
	"log"
)

func main() {

	// fmt.Println("start")
	// defer fmt.Println("middle") // defer means execute this function after main function
	// fmt.Println("end")

	// defer fmt.Println("start")
	// defer fmt.Println("middle")
	// defer fmt.Println("end")
	// output: end, middle, start

	// res, err := http.Get("http://www.google.com/robots.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer res.Body.Close() // most common use of defer
	// // but when working with lots of resources better normally close it
	// robots, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s", robots)

	// a := "start"
	// defer fmt.Println(a)
	// a = "end"
	// it will print "start" beacause defer takes the argumnts when it's called

	// a, b := 1, 0
	// ans := a / b
	// fmt.Println(ans)

	// fmt.Println("start")
	// panic("something bad happened")
	// fmt.Println("end")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello Go!"))
	// })
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// execution queue: exec main -> exec defer -> handle panic -> handle return value
	// fmt.Println("start")
	// defer fmt.Println("this was deferred")
	// panic("something bad happened")
	// fmt.Println("end")

	// fmt.Println("start")
	// defer func() { // anonymous function (doesn't have a name)
	// 	if err := recover(); err != nil { // recover() returns nil if the app isn't panicking
	// 		log.Println("Error: ", err)
	// 	}
	// }() // this prenticies are making that function execute
	// // defer takes function call - not function
	// panic("something bad happened")
	// fmt.Println("end")

	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			panic(err) // rethrow panic if in app this error is something that we can't deal with
			// this will stop execution
		}
	}()
	panic("something bad happened")
	//fmt.Println("done panicking") // this won't execute
}
