package test

import (
	"godis/dataStructure"
	"testing"
)

func TestListNode(t *testing.T) {
	list := dataStructure.ListCreate()
	list.ListAddNodeHead(1)
	list.ListAddNodeTail("hello")
	list.ListAddNodeTail("world")
	if list.ListLength() == 3 {
		t.Log("listAddNodeHead(), ListAddNodeTail(), listNode length()")
	}
	node := list.ListSearchKey("hello")

	list.ListInsertNode(node, "one", true)
	if list.ListLength() == 4 {
		t.Log("listInsertNode()")
	}
	list.ListDelNode(node)
	if list.ListLength() == 3 {
		t.Log("listDelNode()")
	}
	node = list.ListIndex(2)
	list.ListInsertNode(node, 666, true)
	list.ListRotate()
	if list.ListLength() == 4 {
		t.Log("listIndex()")
	}

	list2 := list.ListDup()
	if list2.ListLength() == 4 {
		t.Log("listDup()")
	}

	//head :=
	//list.ListInsertNode()
}
