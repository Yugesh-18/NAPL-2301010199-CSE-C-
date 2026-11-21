package main

import (
	"errors"
	"fmt"
)

type objects struct {
	ID       int
	Name     string
	Quantity int
	Price    int
}

var items []objects

func addingItem(item objects) error {

	if item.ID <= 0 || item.Name == "" || item.Quantity < 0 || item.Price < 0 {
		return errors.New("invalid item details")
	}

	for _, i := range items {
		if i.ID == item.ID {
			return errors.New("item ID already exists")
		}
	}

	items = append(items, item)
	return nil
}

func updatingItem(id int, updated objects) error {
	if updated.Name == "" || updated.Quantity < 0 || updated.Price < 0 {
		return errors.New("invalid item details")
	}

	for i, item := range items {
		if item.ID == id {
			items[i].Name = updated.Name
			items[i].Quantity = updated.Quantity
			items[i].Price = updated.Price
			return nil
		}
	}

	return errors.New("item not found")
}

func removingItem(id int) error {

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			return nil
		}
	}

	return errors.New("item not found")
}

func displayingItems() {

	if len(items) == 0 {
		fmt.Println("Inventory is empty!")
		return
	}

	fmt.Println("\n Inventory Details ")

	for _, item := range items {
		fmt.Println("ID:", item.ID)
		fmt.Println("Name:", item.Name)
		fmt.Println("Quantity:", item.Quantity)
		fmt.Println("Price:", item.Price)
	}
}

type summary struct {
	TotalItems int
	TotalValue int
}

func calculateSummary(data []objects, ch chan summary) {

	var result summary

	for _, item := range data {
		result.TotalItems += item.Quantity
		result.TotalValue += item.Quantity * item.Price
	}

	ch <- result
}

func showSummary() {

	ch := make(chan summary)

	go calculateSummary(items, ch)

	result := <-ch

	fmt.Println("\n Inventory Summary ")
	fmt.Println("Total Products:", len(items))
	fmt.Println("Total Quantity:", result.TotalItems)
	fmt.Println("Total Inventory Value:", result.TotalValue)
}

func main() {

	var choice int

	for {

		fmt.Println("\nINVENTORY MANAGEMENT SYSTEM ")
		fmt.Println("1. Add Item")
		fmt.Println("2. Display Items")
		fmt.Println("3. Update Item")
		fmt.Println("4. Remove Item")
		fmt.Println("5. Inventory Summary")
		fmt.Println("6. Exit")

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			var item objects

			fmt.Print("Enter Item ID: ")
			fmt.Scan(&item.ID)

			fmt.Print("Enter Item Name: ")
			fmt.Scan(&item.Name)

			fmt.Print("Enter Quantity: ")
			fmt.Scan(&item.Quantity)

			fmt.Print("Enter Price: ")
			fmt.Scan(&item.Price)

			err := addingItem(item)

			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Item added successfully!")
			}

		case 2:
			displayingItems()

		case 3:
			var id int
			var updated objects

			fmt.Print("Enter Item ID to update: ")
			fmt.Scan(&id)

			fmt.Print("Enter New Name: ")
			fmt.Scan(&updated.Name)

			fmt.Print("Enter New Quantity: ")
			fmt.Scan(&updated.Quantity)

			fmt.Print("Enter New Price: ")
			fmt.Scan(&updated.Price)

			err := updatingItem(id, updated)

			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Item updated successfully!")
			}

		case 4:
			var id int

			fmt.Print("Enter Item ID to remove: ")
			fmt.Scan(&id)

			err := removingItem(id)

			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Item removed successfully!")
			}

		case 5:
			showSummary()

		case 6:
			fmt.Println("Thank you for using Inventory Management System!")
			return

		default:
			fmt.Println("Invalid choice! Please try again.")
		}
	}
}
