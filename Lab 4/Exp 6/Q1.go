package main

import (
	"fmt"
	"sync"
	"time"
)

func getSquares(limit int, report chan<- string, tracker *sync.WaitGroup) {
	defer tracker.Done()
	fmt.Println("[Active] Worker A: Math Squares")

	time.Sleep(12 * time.Millisecond)
	for x := 1; x <= limit; x++ {
		report <- fmt.Sprintf("Square of %d is %d", x, x*x)
	}

	fmt.Println("[Finished] Worker A: Math Squares")
}

func getCubes(limit int, report chan<- string, tracker *sync.WaitGroup) {
	defer tracker.Done()
	fmt.Println("[Active] Worker B: Math Cubes")

	time.Sleep(8 * time.Millisecond)
	for x := 1; x <= limit; x++ {
		report <- fmt.Sprintf("Cube of %d is %d", x, x*x*x)
	}

	fmt.Println("[Finished] Worker B: Math Cubes")
}

func getFibonacci(limit int, report chan<- string, tracker *sync.WaitGroup) {
	defer tracker.Done()
	fmt.Println("[Active] Worker C: Fibonacci Sequence")

	time.Sleep(15 * time.Millisecond)
	val1, val2 := 0, 1
	for x := 1; x <= limit; x++ {
		report <- fmt.Sprintf("Fibonacci pos %d is %d", x, val1)
		val1, val2 = val2, val1+val2
	}

	fmt.Println("[Finished] Worker C: Fibonacci Sequence")
}

func main() {
	dataStream := make(chan string, 25)
	var controlGroup sync.WaitGroup

	fmt.Println("--- Booting up parallel workers ---")

	controlGroup.Add(3)
	go getSquares(5, dataStream, &controlGroup)
	go getCubes(5, dataStream, &controlGroup)
	go getFibonacci(5, dataStream, &controlGroup)

	go func() {
		controlGroup.Wait()
		close(dataStream)
	}()

	fmt.Println("\n--- Capturing outbound stream ---")
	for update := range dataStream {
		fmt.Println(update)
	}

	fmt.Println("\n--- All operations complete ---")
}
