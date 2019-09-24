package main

import "fmt"

func main() {
	// i++ is a statement, not expression

	// for i := 0; i < 5; i++ { // scoped to loop
	// 	fmt.Println(i)
	// }

	// for i := 0; i < 5; i = i + 2 {
	// 	fmt.Println(i)
	// }

	// for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
	// 	fmt.Println(i, j)
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// 	if i%2 == 0 {
	// 		i /= 2
	// 	} else {
	// 		i = 2*i + 1
	// 	}
	// }

	// i := 0 // scoped to main
	// for ; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// i := 0
	// for i < 5 { // infinity loop
	// 	fmt.Println(i)
	// 	// i++ // no more inf
	// }

	// i := 0
	// for {
	// 	fmt.Println(i)
	// 	i++
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// prints odd numbers (even is an opposite)
	// for i := 0; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// for i := 1; i <= 3; i++ {
	// 	for j := 1; j <= 3; j++ {
	// 		fmt.Println(i * j)
	// 		if i * j >= 3 {
	// 			break
	// 		}
	// 	}
	// }

	// Loop:
	// 	for i := 1; i <= 3; i++ {
	// 		for j := 1; j <= 3; j++ {
	// 			fmt.Println(i * j)
	// 			if i*j >= 3 {
	// 				break Loop
	// 			}
	// 		}
	// 	}

	// s := []int{1, 2, 3}
	// fmt.Println(s)
	// for k, v := range s { // key, value
	// 	fmt.Println(k, v)
	// }

	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	for k, v := range statePopulations {
		fmt.Println(k, v)
	}

	for k := range statePopulations {
		fmt.Println(k)
	}

	for _, v := range statePopulations {
		fmt.Println(v)
	}

	// s := "hello Go!"
	// for k, v := range s {
	// 	fmt.Println(k, string(v))
	// }
}
