package main

import (
	"fmt"
	"time"
)

// channels are really designed to synchronise
// we can create channel only by "make"
// data transmission between multiple goroutines
// var wg = sync.WaitGroup{}

// func main() {
// 	ch := make(chan int)
// 	wg.Add(2)
// 	go func() { // receiving goroutine
// 		i := <-ch
// 		fmt.Println(i)
// 		wg.Done()
// 	}()
// 	go func() { // sending goroutine
// 		i := 42
// 		ch <- i
// 		i = 27
// 		wg.Done()
// 	}()
// 	wg.Wait()
// }

// var wg = sync.WaitGroup{}

// func main() {
// 	ch := make(chan int)
// 	for j := 0; j < 5; j++ {
// 		wg.Add(2)
// 		go func() { // putting this in another scope will cause error -> example below
// 			i := <-ch
// 			fmt.Println(i)
// 			wg.Done()
// 		}()
// 		go func() {
// 			ch <- 42
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// }

// var wg = sync.WaitGroup{}

// func main() {
// 	ch := make(chan int)
// 	go func() { // putting this in this scope will cause error (deadlock)
// 		i := <-ch
// 		fmt.Println(i)
// 		wg.Done()
// 	}()
// 	for j := 0; j < 5; j++ {
// 		wg.Add(2)
// 		go func() {
// 			ch <- 42  // this gonna pause the execution of this goroutine right at this line until there is a space available in the channel
// 			wg.Done() // by default we're working with unbuffered channels which means only one message can be in the channel at one time!!!
// 		}()
// 	}
// 	wg.Wait()
// }

// var wg = sync.WaitGroup{}

// func main() {
// 	ch := make(chan int)
// 	wg.Add(2)
// 	go func() {
// 		i := <-ch
// 		fmt.Println(i)
// 		ch <- 27
// 		wg.Done()
// 	}()
// 	go func() {
// 		ch <- 42
// 		fmt.Println(<-ch)
// 		wg.Done()
// 	}()
// 	wg.Wait()
// }

// var wg = sync.WaitGroup{}

// func main() {
// 	ch := make(chan int)
// 	wg.Add(2)
// 	go func(ch <-chan int) { // receive only channel
// 		i := <-ch
// 		fmt.Println(i)
// 		// ch <- 27 // we sand send data in receive only channel
// 		wg.Done()
// 	}(ch)
// 	go func(ch chan<- int) { // send olny channel
// 		ch <- 42
// 		// fmt.Println(<-ch) // we can't recaive data from send only channel
// 		wg.Done()
// 	}(ch)
// 	wg.Wait()
// }

// var wg = sync.WaitGroup{}

// func main() {
// 	ch := make(chan int, 50) // adding a buffer that can store 50 integers (for assymetric loading)
// 	wg.Add(2)
// 	go func(ch <-chan int) { // receive only channel
// 		i := <-ch
// 		fmt.Println(i)
// 		i = <-ch
// 		fmt.Println(i)
// 		wg.Done()
// 	}(ch)
// 	go func(ch chan<- int) { // send olny channel
// 		ch <- 42
// 		ch <- 27 // deadlock on this line
// 		wg.Done()
// 	}(ch)
// 	wg.Wait()
// }

// var wg = sync.WaitGroup{}

// func main() {
// 	ch := make(chan int, 50)
// 	wg.Add(2)
// 	go func(ch <-chan int) {
// 		for i := range ch {
// 			fmt.Println(i)
// 		}
// 		wg.Done()
// 	}(ch)
// 	go func(ch chan<- int) {
// 		ch <- 42
// 		// close(ch) // don't do it if u still sending data; you even can't know that channel is closed
// 		ch <- 27
// 		close(ch) // to prevent deadlock in recaive only goroutune (for loop)
// 		wg.Done()
// 	}(ch)
// 	wg.Wait()
// }

// we can use ok syntax to get to know if the channel is closed
// var wg = sync.WaitGroup{}

// func main() {
// 	ch := make(chan int, 50)
// 	wg.Add(2)
// 	go func(ch <-chan int) {
// 		for {
// 			if i, ok := <-ch; ok { // if channel is open then ok is true
// 				fmt.Println(i)
// 			} else {
// 				break
// 			}
// 		} // this is functionally exactly the same as the for range construct
// 		wg.Done()
// 	}(ch)
// 	go func(ch chan<- int) {
// 		ch <- 42
// 		ch <- 27
// 		close(ch)
// 		wg.Done()
// 	}(ch)
// 	wg.Wait()
// }

// const (
// 	logInfo    = "INFO"
// 	logWarning = "WARNING"
// 	logError   = "ERROR"
// )

// type logEntry struct {
// 	time     time.Time
// 	severity string
// 	message  string
// }

// var logCh = make(chan logEntry, 50)

// func main() {
// 	go logger()
// 	// defer func() {
// 	// 	close(logCh)
// 	// }()
// 	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

// 	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
// 	time.Sleep(100 + time.Millisecond)
// }

// func logger() {
// 	for entry := range logCh {
// 		fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
// 	}
// }

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // signal only channel
// struct{} needs zero memory allocation

func main() {
	go logger()
	logCh <- logEntry{time.Now(), logInfo, "App is starting"}

	logCh <- logEntry{time.Now(), logInfo, "App is shutting down"}
	time.Sleep(100 + time.Millisecond)
	doneCh <- struct{}{} // defining a struct with no fields and initializing that struct using this curly brackets {}
}

func logger() {
	for {
		select { // select cause entire statement is gonna block until a message is received on one of the channles that is listening for
		case entry := <-logCh:
			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			break
		} // default is used for non blocking select statement
	}
}

// default
// if there is a message ready on one of the channels that are being monitored
// then it's gonna execute that code path
// if not - it will execute the default block
