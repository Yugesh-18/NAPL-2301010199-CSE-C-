package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	switch num {
	case 1:
		fmt.Println("Performing operations for int")
	case 2:
		fmt.Println("Performing operations for float")
	default:
		fmt.Println("Exiting the program")
	}
}
