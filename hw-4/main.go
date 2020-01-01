package main

type Item struct {
	value interface{}
	next *Item
	prev *Item
}

func (item *Item) Value() interface{} {
	return item.value
}

func (item *Item) Next() *Item {
	return item.next
}

func (item *Item) Prev() *Item {
	return item.prev
}

type List struct {
	head *Item
	tail *Item
	size int
}

func (list *List) Len() int {
	return list.size
}

func (list *List) First() *Item {
	return list.head
}

func (list *List) Last() *Item {
	return list.tail
}

func (list *List) PushFront(v interface{}) *Item {
	itemPtr := &Item{
		value: v,
	}

	head := list.First()
	tail := list.Last()

	if head == nil || tail == nil {
		list.head = itemPtr
		list.tail = itemPtr
	} else {
		itemPtr.next = head
		list.head.prev = itemPtr
		list.head = itemPtr
	}

	list.size++

	return itemPtr
}

func (list *List) PushBack(v interface{}) *Item {
	itemPtr := &Item {
		value: v,
	}

	if list.First() == nil || list.Last() == nil {
		list.head = itemPtr
		list.tail = itemPtr
	} else {
		itemPtr.prev = list.Last()
		list.Last().next = itemPtr
		list.tail = itemPtr
	}

	list.size++

	return itemPtr
}

func (list *List) Remove(item *Item) {
	if item.prev != nil {
		item.prev.next = item.next
	}

	if item.next != nil {
		item.next.prev = item.prev
	}

	if item == list.head {
		list.head = item.next
	}

	if item == list.tail {
		list.tail = item.prev
	}

	item.next = nil
	item.prev = nil

	list.size--
}

func main() {}