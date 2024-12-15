// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var name string
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	fmt.Print("Enter your name: ")
// 	fmt.Scanln(&name)
// 	go hello(name, &wg)
// 	go bye(name, &wg)
// 	wg.Wait()
// }

// func hello(name string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Hello %s!\n", name)
// }

// func bye(name string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Good Bye %s!\n", name)
// }

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	msg := make(chan string)

	// Add one to the wait group since we are starting one goroutine
	wg.Add(1)

	// Start the greet goroutine
	go greet(msg, &wg)

	// Wait for the greet goroutine to finish
	greeting := <-msg // This will now receive after the greet goroutine sends the message

	// Ensure that the greet goroutine is done before printing the result
	wg.Wait()

	fmt.Println("Greeting received")
	fmt.Println(greeting)

	_, ok := <-msg
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}
}

func greet(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrease the counter when greet finishes

	fmt.Println("Greeter waiting to send greeting!")

	// Simulate some work
	ch <- "Hello Krishna" // Send the greeting to the channel
	close(ch)
	fmt.Println("Greeter completed")
}
