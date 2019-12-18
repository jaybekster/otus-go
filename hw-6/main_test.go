package main

import (
	"testing"
)

func TestList(t *testing.T) {
	list := List{}

	if list.Len() != 0 {
	   t.Errorf("Len is working incorrect with empty")
	}

	middleItem := list.PushBack("last item")
	firstItem := list.PushFront("middle item")
	lastItem := list.PushBack("first item")

	if list.Len() != 3 {
		t.Errorf("Len is working incorrect after PushBack or PushFront methods")
	}

	if list.First() != firstItem {
		t.Errorf("Head Item is not correct")
	}

	if list.Last() != lastItem {
		t.Errorf("Tail Item is not correct")
	}

	if (*list.First()).Next() != middleItem {
		t.Errorf("Middle Item is not correct")
	}

	list.Remove(list.First())

	if list.Len() != 2 {
		t.Errorf("Len is working incorrect after Remove method")
	}

	if (*list.First()).Next() != list.Last() {
		t.Errorf("Middle Item has not been removed correctly")
	}
}
