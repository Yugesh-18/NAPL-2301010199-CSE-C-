package main

import "fmt"

func main() {

	// Part 1: Slice Operations

	var n int
	fmt.Print("Enter number of students: ")
	fmt.Scan(&n)

	students := make([]string, n)

	// Taking student names as input
	for i := 0; i < n; i++ {
		fmt.Printf("Enter student %d name: ", i+1)
		fmt.Scan(&students[i])
	}

	fmt.Println("\nInitial Slice:", students)

	// Adding a student
	var newStudent string
	fmt.Print("\nEnter student name to add: ")
	fmt.Scan(&newStudent)

	students = append(students, newStudent)
	fmt.Println("After Adding:", students)

	// Removing student by index
	var index int
	fmt.Print("\nEnter index to remove: ")
	fmt.Scan(&index)

	if index >= 0 && index < len(students) {
		students = append(students[:index], students[index+1:]...)
		fmt.Println("After Removing:", students)
	} else {
		fmt.Println("Invalid index!")
	}

	// Updating student
	fmt.Print("\nEnter index to update: ")
	fmt.Scan(&index)

	if index >= 0 && index < len(students) {
		var updatedName string
		fmt.Print("Enter new student name: ")
		fmt.Scan(&updatedName)

		students[index] = updatedName
		fmt.Println("After Updating:", students)
	} else {
		fmt.Println("Invalid index!")
	}

	// Part 2: Map Operations

	marks := make(map[string]int)

	var subjects int
	fmt.Print("\nEnter number of subjects: ")
	fmt.Scan(&subjects)

	// Inserting subject and marks
	for i := 0; i < subjects; i++ {
		var subject string
		var mark int

		fmt.Printf("Enter subject %d: ", i+1)
		fmt.Scan(&subject)

		fmt.Printf("Enter marks for %s: ", subject)
		fmt.Scan(&mark)

		marks[subject] = mark
	}

	fmt.Println("\nInitial Map:", marks)

	// Inserting another subject
	var subject string
	var mark int

	fmt.Print("\nEnter subject to insert: ")
	fmt.Scan(&subject)

	fmt.Print("Enter marks: ")
	fmt.Scan(&mark)

	marks[subject] = mark
	fmt.Println("After Inserting:", marks)

	// Deleting a subject
	fmt.Print("\nEnter subject to delete: ")
	fmt.Scan(&subject)

	if _, exists := marks[subject]; exists {
		delete(marks, subject)
		fmt.Println("After Deleting:", marks)
	} else {
		fmt.Println("Subject not found!")
	}

	// Looking up a subject
	fmt.Print("\nEnter subject to lookup: ")
	fmt.Scan(&subject)

	if mark, exists := marks[subject]; exists {
		fmt.Println("Marks:", mark)
	} else {
		fmt.Println("Subject not found!")
	}
}
