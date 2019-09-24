package main

// import (
// 	"fmt"
// 	"reflect"
// )

// type Doctor struct {
// 	number     int
// 	actorName  string
// 	companions []string
// 	episodes   []string
// }

// type Doctor struct {
// 	Number     int
// 	ActorName  string
// 	Companions []string
// 	Episodes   []string
// }

// type Animal struct {
// 	Name   string `required max:"100"`
// 	Origin string
// }

// type Bird struct {
// 	Animal   // works like inheritance (embedding)
// 	SpeedKPH float32
// 	CanFly   bool
// }

func main() {
	// Map

	// statePopulations := make(map[string]int)
	// statePopulations = map[string]int{
	// 	"California":   39250017,
	// 	"Texas":        27862596,
	// 	"Florida":      20612439,
	// 	"New York":     19745289,
	// 	"Pennsylvania": 12802503,
	// 	"Illinois":     12801539,
	// 	"Ohio":         11614373,
	// }
	// m := map[[3]int]string{}
	// fmt.Println(statePopulations, m)
	// fmt.Println(statePopulations["Ohio"])

	// statePopulations["Georgia"] = 10310371
	// fmt.Println(statePopulations["Georgia"])
	// fmt.Println(statePopulations)

	// delete(statePopulations, "Georgia")
	// fmt.Println(statePopulations)
	// fmt.Println(statePopulations["Georgia"])

	// // pop, ok := statePopulations["Oho"]
	// // fmt.Println(pop, ok)

	// _, ok := statePopulations["Ohio"]
	// fmt.Println(ok)
	// // it don't have be "ok" word, but it is commonly used for this kind of test
	// _, ulala := statePopulations["Ohio"]
	// fmt.Println(ulala)

	// fmt.Println(len(statePopulations))

	// sp := statePopulations // referenction
	// delete(sp, "Ohio")
	// fmt.Println(len(sp))
	// fmt.Println(len(statePopulations))

	// Struct

	// aDoctor := Doctor{
	// 	Number:    3,
	// 	ActorName: "Jon Pertwee",
	// 	Companions: []string{
	// 		"Liz Shaw",
	// 		"Jo Grant",
	// 		"Sarah Jane Smith",
	// 	},
	// }
	// fmt.Println(aDoctor.companions[1])

	// aDoctor := Doctor{
	// 	3,
	// 	"Jon Pertwee",
	// 	[]string{
	// 		"Liz Shaw",
	// 		"Jo Grant",
	// 		"Sarah Jane Smith",
	// 	},
	// }
	// it is valid go syntax but better don't use it
	// because if u will expand the struct those fields
	// could be initialized with invalid data
	//fmt.Println(aDoctor)

	// // anonymous struct
	// // common use cases
	// // very short struct life such as
	// // generating a json response to web service call
	// // only used one time
	// aDoctor := struct{ name string }{name: "John Pertwee"}
	// // anotherDoctor := aDoctor // copy
	// anotherDoctor := &aDoctor // original
	// anotherDoctor.name = "Tom Baker"
	// fmt.Println(aDoctor)
	// fmt.Println(anotherDoctor)

	// b := Bird{}
	// b.Name = "Emu"
	// b.Origin = "Australia"
	// b.SpeedKPH = 48
	// b.CanFly = false
	// fmt.Println(b)
	// fmt.Println(b.Name)

	// b := Bird{
	// 	Animal:   Animal{Name: "Emu", Origin: "Australia"},
	// 	SpeedKPH: 48,
	// 	CanFly:   false,
	// }
	// fmt.Println(b.Name)

	// t := reflect.TypeOf(Animal{})
	// field, _ := t.FieldByName("Name")
	// fmt.Println(field.Tag)
}
