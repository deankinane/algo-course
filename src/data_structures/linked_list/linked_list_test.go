package datastructures

import (
	"fmt"
	"testing"
)

type Test struct {
	name string
	val  int
}

func getTestList() LinkedList[Test] {
	list := LinkedList[Test]{}

	for i := 0; i < 10; i++ {
		list.Append(Test{
			name: fmt.Sprint("Item", i),
			val:  i,
		})
	}
	return list
}

func getTestItem() Test {
	return Test{
		name: "ligma",
		val:  69,
	}
}

func TestLinkedList_AppendShouldAddToEndOfList(t *testing.T) {
	list := getTestList()

	if list.Head.Item.name != "Item0" || list.Head.Item.val != 0 {
		t.Errorf("Exptected Item0, got %v", list.Head.Item.name)
	}

	if list.Tail.Item.name != "Item9" || list.Tail.Item.val != 9 {
		t.Errorf("Exptected Item9, got %v", list.Head.Item.name)
	}
}

func TestLinkedList_PrependShouldAddToStartOfList(t *testing.T) {
	list := getTestList()

	list.Prepend(getTestItem())

	if list.Head.Item.name != "ligma" {
		t.Errorf("Expected ligma, got %v", list.Head.Item.name)
	}
}

func TestLinkedList_InsertAtShoudlAddAtCorrectIndex(t *testing.T) {
	list := getTestList()

	list.InsertAt(getTestItem(), 5)

	cur := list.Head
	for i := 0; i < 5; i++ {
		cur = cur.Next
	}

	if cur.Item.name != "ligma" {
		t.Errorf("Expected ligma, got %v", cur.Item.name)
	}
}

func TestLinkedList_RemoveShouldRemoveItemFromTheList(t *testing.T) {
	list := getTestList()

	cur := list.Head
	for i := 0; i < 5; i++ {
		cur = cur.Next
	}

	removed, err := list.Remove(cur.Item)
	if err != nil {
		t.Errorf("Expected %v, got error: %v", cur.Item.name, err)
	} else if removed.name != cur.Item.name {
		t.Errorf("Expected %v, got %v", cur.Item.name, removed.name)
	}

	cur = list.Head
	for cur != nil {
		if cur.Item.name == removed.name {
			t.Errorf("Expected not to find item but found item")
			break
		}
		cur = cur.Next
	}
}

func TestLinkedList_RemoveAtShouldRemoveItemAtIndex(t *testing.T) {
	list := getTestList()

	cur := list.Head
	for i := 0; i < 6; i++ {
		cur = cur.Next
	}

	removed, err := list.RemoveAt(6)
	if err != nil {
		t.Errorf("Expected %v, got error: %v", cur.Item.name, err)
	} else if removed.name != cur.Item.name {
		t.Errorf("Expected %v, got %v", cur.Item.name, removed.name)
	}

	cur = list.Head
	for cur != nil {
		if cur.Item.name == removed.name {
			t.Errorf("Expected not to find item but found item")
			break
		}
		cur = cur.Next
	}
}

func TestLinkedList_GetShouldReturnItemAtIndex(t *testing.T) {
	list := getTestList()

	if item, _ := list.Get(5); item.name != "Item5" {
		t.Errorf("Expected Item5, got %v", item.name)
	}

	if item, _ := list.Get(8); item.name != "Item8" {
		t.Errorf("Expected Item8, got %v", item.name)
	}

	if item, err := list.Get(12); err == nil {
		t.Errorf("Expected error, got %v", item)
	}
}
