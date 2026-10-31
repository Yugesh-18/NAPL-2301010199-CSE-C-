package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Job    string
	Salary float64
}

func (p *Person) ReadData() {
	fmt.Print("Enter Name: ")
	fmt.Scanln(&p.Name)

	fmt.Print("Enter Age: ")
	fmt.Scanln(&p.Age)

	fmt.Print("Enter Job: ")
	fmt.Scanln(&p.Job)

	fmt.Print("Enter Salary: ")
	fmt.Scanln(&p.Salary)
}

func (p Person) DisplayData() {
	fmt.Println("Name   :", p.Name)
	fmt.Println("Age    :", p.Age)
	fmt.Println("Job    :", p.Job)
	fmt.Printf("Salary : %.2f\n", p.Salary)
}

func main() {
	var person1 Person
	var person2 Person

	fmt.Println("Enter details for Person 1:")
	person1.ReadData()

	fmt.Println("\nEnter details for Person 2:")
	person2.ReadData()

	fmt.Println("\n===== Person 1 =====")
	person1.DisplayData()

	fmt.Println("\n===== Person 2 =====")
	person2.DisplayData()
}
