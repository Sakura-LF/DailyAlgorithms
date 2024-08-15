package linkedlist

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	//myLinkedList := Constructor()
	//myLinkedList.AddAtHead(1)
	//myLinkedList.AddAtHead(2)
	//myLinkedList.AddAtTail(4)
	//myLinkedList.AddAtTail(7)
	//myLinkedList.Print()
	//
	////fmt.Println(myLinkedList.Get(3))
	////fmt.Println(myLinkedList.Size)
	//// Delete
	//fmt.Println(myLinkedList.Get(4))
	//myLinkedList.DeleteAtIndex(2)
	//myLinkedList.Print()
}

func TestConstructor1(t *testing.T) {
	myLinkedList := Constructor()
	myLinkedList.AddAtHead(7)
	myLinkedList.AddAtHead(2)
	myLinkedList.AddAtHead(1)
	myLinkedList.AddAtIndex(3, 0)
	myLinkedList.DeleteAtIndex(2)
	myLinkedList.AddAtHead(6)
	myLinkedList.AddAtTail(4)
	myLinkedList.Print()
	myLinkedList.DeleteAtIndex(4)
	myLinkedList.Print()
	//fmt.Println(myLinkedList.Get(4))

}
func TestConstructor2(t *testing.T) {
	myLinkedList := Constructor()
	myLinkedList.AddAtHead(2)
	myLinkedList.DeleteAtIndex(1)
}
