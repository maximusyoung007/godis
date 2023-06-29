package dataStructure

import "reflect"

const AL_START_HEAD int = 0
const AL_START_TAIL int = 1

type ListNode[T any] struct {
	pre  *ListNode[T]
	next *ListNode[T]
	val  T
}

type List[T any] struct {
	head   *ListNode[T]
	tail   *ListNode[T]
	length int
}

type ListIter[T any] struct {
	next *ListNode[T]

	direction int
}

func ListCreate() *List[any] {
	var list *List[any]
	list = &List[any]{nil, nil, 0}
	return list
}

// ListGetIterator 创建一个迭代器，调用listNext()方法返回链表的下一个节点
func (list *List[any]) ListGetIterator(direction int) *ListIter[any] {
	var iter *ListIter[any]
	iter = &ListIter[any]{nil, direction}
	if direction == AL_START_HEAD {
		iter.next = list.head
	} else if direction == AL_START_TAIL {
		iter.next = list.tail
	} else {
		return nil
	}
	//iter.direction = direction
	return iter
}

func (iter *ListIter[T]) ListNext() *ListNode[T] {
	//生成迭代器的时候，iter.next = list.head或者iter.next = list.tail，所以current为head或者tail
	current := iter.next
	if current != nil {
		if iter.direction == AL_START_HEAD {
			//向后迭代
			iter.next = current.next
		} else {
			//向前迭代
			iter.next = current.pre
		}
	}
	return current
}

func (list *List[T]) ListLength() int {
	return list.length
}

func (list *List[T]) ListFirst() *ListNode[T] {
	return list.head
}

func (list *List[T]) ListLast() *ListNode[T] {
	return list.tail
}

func (listNode *ListNode[T]) listPreNode() *ListNode[T] {
	return listNode.pre
}

func (listNode *ListNode[T]) listNextNode() *ListNode[T] {
	return listNode.next
}

func (listNode *ListNode[T]) listNodeValue() T {
	return listNode.val
}

// ListAddNodeHead 将一个新节点加入到链表表头
func (list *List[T]) ListAddNodeHead(value T) {
	var node *ListNode[T]
	node = &ListNode[T]{nil, nil, value}
	if list.length == 0 {
		list.head = node
		list.tail = node
		node.pre = nil
		node.next = nil
	} else {
		node.pre = nil
		node.next = list.head
		list.head.pre = node
		list.head = node
	}
	list.length++
}

// ListAddNodeTail 将新节点添加到链表表尾
func (list *List[T]) ListAddNodeTail(value T) {
	var node *ListNode[T]
	node = &ListNode[T]{nil, nil, value}
	node.val = value
	if list.length == 0 {
		list.head = node
		list.tail = node
		node.pre = nil
		node.next = nil
	} else {
		node.next = nil
		node.pre = list.tail
		list.tail.next = node
		list.tail = node
	}
	list.length++
}

// ListInsertNode 将节点添加到给定节点之前或者之后, after==false之前，after==true之后
func (list *List[T]) ListInsertNode(oldNode *ListNode[T], value T, after bool) {
	var node *ListNode[T]
	node = &ListNode[T]{nil, nil, value}

	if after {
		node.pre = oldNode
		node.next = oldNode.next
		if oldNode == list.tail {
			list.tail = node
		}
	} else {
		node.next = oldNode
		node.pre = oldNode.pre
		if oldNode == list.head {
			list.head = node
		}
	}

	if node.pre != nil {
		node.pre.next = node
	}
	if node.next != nil {
		node.next.pre = node
	}
	list.length++
}

// ListDelNode 删除节点
func (list *List[T]) ListDelNode(node *ListNode[T]) {
	if node.pre != nil {
		//非头节点
		node.pre.next = node.next
	} else {
		//头节点
		list.head = node.next
	}

	if node.next != nil {
		node.next.pre = node.pre
	} else {
		list.tail = node.pre
	}
	//list.free(node.val)
	list.length--
}

// ListSearchKey 返回给定值的节点
func (list *List[T]) ListSearchKey(key T) *ListNode[T] {
	iter := list.ListGetIterator(AL_START_HEAD)
	node := iter.ListNext()
	for node != nil {
		if reflect.DeepEqual(key, node.val) {
			return node
		}
		node = iter.ListNext()
	}
	return nil
}

// ListIndex 返回给定索引的节点,从0开始，负数表示从链表末尾开始，从-1开始
func (list *List[T]) ListIndex(index int) *ListNode[T] {
	var node *ListNode[T]
	if index < 0 {
		index = (-index) - 1
		node = list.tail
		for index >= 0 && node.pre != nil {
			index--
			node = node.pre
		}
	} else {
		node = list.head
		for index >= 0 && node.next != nil {
			index--
			node = node.next
		}
	}
	return node
}

// ListRotate 取出链表尾节点并插入表头
func (list *List[T]) ListRotate() {
	tail := list.tail
	if list.ListLength() <= 1 {
		return
	}
	list.tail = tail.pre
	list.tail.next = nil

	list.head.pre = tail
	tail.pre = nil
	tail.next = list.head
	list.head = tail
}

// ListDup 复制一个给定链表
func (list *List[T]) ListDup() *List[any] {
	copy := ListCreate()
	var iter *ListIter[T]
	var node *ListNode[T]

	//copy.dup = list.dup
	//copy.free = list.free
	//copy.match = list.match

	iter = list.ListGetIterator(AL_START_HEAD)
	node = iter.ListNext()
	for node != nil {
		var value T
		value = node.val
		copy.ListAddNodeTail(value)
		node = iter.ListNext()
	}
	return copy
}
