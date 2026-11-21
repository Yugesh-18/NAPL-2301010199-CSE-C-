package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Grade string
}

func modifyVariable(ptr *int) {
	*ptr = 11
}

func addDataInStructure(s *Student) {
	s.Name = "Navi"
	s.Age = 21
	s.Grade = "C"
}

func main() {
	var myVar int = 4
	var ptr *int = &myVar
	fmt.Printf("1. Address of the variable using pointer: %p\n", ptr)
	fmt.Printf("2. Value of the variable using pointer: %d\n", *ptr)
	fmt.Printf("3. Value of the variable before modification: %d\n", myVar)
	modifyVariable(&myVar)
	fmt.Printf("4. Value of the variable after modification: %d\n\n", myVar)
	s1 := new(Student)
	addDataInStructure(s1)
	fmt.Println("---> Student Structure Data <---")
	fmt.Printf("Name: %s\n", s1.Name)
	fmt.Printf("Age: %d\n", s1.Age)
	fmt.Printf("Grade: %s\n", s1.Grade)

}
