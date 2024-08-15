package another

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	constructor := Constructor()
	constructor.AddAtHead(7)
	constructor.AddAtHead(2)
	constructor.AddAtHead(1)
	constructor.AddAtIndex(3, 0)

	constructor.DeleteAtIndex(2)
	constructor.AddAtHead(6)
	constructor.AddAtTail(4)
	constructor.printLinkedList()
	//fmt.Println(constructor.Get(4))

}
