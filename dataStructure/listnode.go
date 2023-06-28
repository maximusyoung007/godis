package dataStructure

type ListNode[T any] struct {
	pre  *ListNode[T]
	next *ListNode[T]
	val  T
}

type List[T any] struct {
	head   *ListNode[T]
	tail   *ListNode[T]
	length int

	/**
	不知道怎么用 先放着
	// 节点值复制函数
	    void *(*dup)(void *ptr);

	    // 节点值释放函数
	    void (*free)(void *ptr);

	    // 节点值对比函数
	    int (*match)(void *ptr, void *key);
	*/
}
