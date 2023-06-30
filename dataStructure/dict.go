package dataStructure

// DictEntry 哈希表节点
type DictEntry[T any] struct {
	key   T
	value T
	//将多个哈希值相同的键值对放在一起
	next *DictEntry[T]
}

// DictHt 哈希表
type DictHt[T any] struct {
	table []*DictEntry[T]
	size  int
	//哈希表大小掩码，计算键应该放到table数组的哪个索引上
	sizeMask int
	used     int
}

type Dict[T any] struct {
	//两个数组，一般情况下只使用h[0]，h[1]在rehash时使用
	ht [2]DictHt[T]

	//rehash索引，不进行rehash时，值为-1
	rehashIndex int

	//目前正在运行的安全迭代器的数量
	iterators int
}
