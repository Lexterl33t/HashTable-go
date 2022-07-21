package main

import "fmt"

const size = 7

type Hashtable struct {
	list [size]*List
}

type List struct {
	Head *Node
}

type Node struct {
	Key  string
	Next *Node
}

func InitHash() *Hashtable {
	hashtable := Hashtable{}
	for i := 0; i < size; i++ {
		hashtable.list[i] = &List{}
	}
	return &hashtable
}

func resolve_hash(key string) int {

	sum := 0
	for _, b := range key {
		sum += int(b)
	}

	return sum % size
}

func (h *Hashtable) Insert(key string) {
	if _, err := h.Search(key); err != nil {
		index := resolve_hash(key)
		h.list[index].insert(key)

	} else {
		fmt.Println("Key already taken")
	}
}

func (h *Hashtable) Search(key string) (*Node, error) {
	index := resolve_hash(key)
	if node, err := h.list[index].search(key); err != nil {
		return nil, err
	} else {
		return node, nil
	}
}

func (l *List) search(key string) (*Node, error) {
	actual_node := l.Head
	for actual_node != nil {
		if actual_node.Key == key {
			return actual_node, nil
		}
		actual_node = actual_node.Next
	}
	return nil, fmt.Errorf("key not found")
}

func (l *List) insert(key string) {
	node := Node{
		Key:  key,
		Next: l.Head,
	}

	l.Head = &node
}

func (l *List) PrintLinkedList() {
	actual := l.Head
	for actual != nil {
		fmt.Printf("%s, ", actual.Key)
		actual = actual.Next
	}
}

func (h *Hashtable) PrintHashtable() {
	for i, d := range h.list {
		fmt.Printf("%d => ", i)
		d.PrintLinkedList()
		fmt.Printf("\n")
	}
}

func main() {
	table := InitHash()
	table.Insert("Bryton")
	table.Insert("Brytonette")
	table.Insert("Faycan")

	table.PrintHashtable()
}
