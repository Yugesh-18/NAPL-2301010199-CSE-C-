package main

import (
	"App/mathutil"
	"fmt"
)

func main() {
	fmt.Println("Enter two numbers to add:")
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Print("The sum of", a, "and", b, "is:")
	fmt.Println(mathutil.Add(a, b))
	fmt.Println("Enter a number to calculate its factorial:")
	var n int
	fmt.Scan(&n)
	fmt.Print("The factorial of", n, "is:")
	fmt.Println(mathutil.Fact(n))
	fmt.Println("Enter a base and an exponent to calculate power:")
	var base, exponent int
	fmt.Scan(&base, &exponent)
	fmt.Print(base, "raised to the power of", exponent, "is:")
	fmt.Println(mathutil.Pow(base, exponent))
	fmt.Println("Enter a string to reverse:")
	var str string
	fmt.Scan(&str)
	fmt.Println("The reverse of", str, "is:")
	fmt.Println(mathutil.Reverse(str))
	fmt.Println("Enter a string to count vowels:")
	fmt.Scan(&str)
	fmt.Println("The number of vowels in", str, "is:")
	fmt.Println(mathutil.VowelCount(str))
}
