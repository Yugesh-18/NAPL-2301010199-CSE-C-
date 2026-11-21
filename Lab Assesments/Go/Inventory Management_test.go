package main

import "testing"

func TestAddingItem(t *testing.T) {
	items = nil

	item := objects{
		ID:       1,
		Name:     "Laptop",
		Quantity: 5,
		Price:    40000,
	}

	err := addingItem(item)

	if err != nil {
		t.Fatal("Item was not added:", err)
	}

	if len(items) != 1 {
		t.Fatal("Item count is incorrect")
	}
}

func TestUpdatingItem(t *testing.T) {
	items = []objects{
		{1, "Laptop", 5, 40000},
	}

	updated := objects{
		Name:     "Dell Laptop",
		Quantity: 10,
		Price:    45000,
	}

	err := updatingItem(1, updated)

	if err != nil {
		t.Fatal("Item was not updated:", err)
	}

	if items[0].Name != "Dell Laptop" {
		t.Fatal("Item name was not updated")
	}
}

func TestRemovingItem(t *testing.T) {
	items = []objects{
		{1, "Laptop", 5, 40000},
	}

	err := removingItem(1)

	if err != nil {
		t.Fatal("Item was not removed:", err)
	}

	if len(items) != 0 {
		t.Fatal("Item still exists")
	}
}
