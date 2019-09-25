package main

import (
	"fmt"
	"time"
)

// 1
// func main() {
// 	go sayHello()
// 	time.Sleep(100 * time.Millisecond)
// }

// func sayHello() {
// 	fmt.Println("Hello")
// }

// 2
// func main() {
// 	var msg = "Hello"
// 	go func() { // not a good idea
// 		fmt.Println(msg)
// 	}()
// 	msg = "Goodbye"
// 	time.Sleep(100 * time.Millisecond)
// } // printed Goodbye

// 3
// func main() {
// 	var msg = "Hello"
// 	go func(msg string) {
// 		fmt.Println(msg)
// 	}(msg)
// 	msg = "Goodbye"
// 	time.Sleep(100 * time.Millisecond)
// } // printed Hello

// 4
// var wg = sync.WaitGroup{}

// func main() {
// 	var msg = "Hello"
// 	wg.Add(1)
// 	go func(msg string) {
// 		fmt.Println(msg)
// 		wg.Done()
// 	}(msg)
// 	msg = "Goodbye"
// 	wg.Wait()
// }

// 5 \/ this won't work properly
// var wg = sync.WaitGroup{}
// var counter = 0

// func main() {
// 	for i := 0; i < 10; i++ {
// 		wg.Add(2)
// 		go sayHello()
// 		go increment()
// 	}
// 	wg.Wait()
// }

// func sayHello() {
// 	fmt.Printf("Hello #%v\n", counter)
// 	wg.Done()
// }

// func increment() {
// 	counter++
// 	wg.Done()
// }

// 6
// var wg = sync.WaitGroup{}
// var counter = 0
// var m = sync.RWMutex{} // Read Write Mutex - a lock that the application is gonna honour

// func main() {
// 	runtime.GOMAXPROCS(100)
// 	for i := 0; i < 10; i++ {
// 		wg.Add(2)
// 		m.RLock()
// 		go sayHello()
// 		m.Lock()
// 		go increment()
// 	}
// 	wg.Wait()
// }

// func sayHello() {
// 	fmt.Printf("Hello #%v\n", counter)
// 	m.RUnlock()
// 	wg.Done()
// }

// func increment() {
// 	counter++
// 	m.Unlock()
// 	wg.Done()
// }

// 7
// func main() {
// 	runtime.GOMAXPROCS(1) // need to tune this value for app
// 	fmt.Println("Threads: %v", runtime.GOMAXPROCS(-1))
// }

// 8
func main() {
	var msg = "Hello"
	go func() { // not a good idea
		fmt.Println(msg)
	}()
	msg = "Goodbye"
	time.Sleep(100 * time.Millisecond)
} // clear;go run -race src/github.com/kamilcinner/go-goroutines/Main.go
// -race
