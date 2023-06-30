package main

import (
	"fmt"
	"unsafe"
)

type T struct {
	a int
	b int
	r []rune
}

//type F struct {
//	c int
//	d int
//}

func t2(t *[]rune) *T {
	//var dummy T
	var x struct{}
	dummy := (*T)(unsafe.Pointer(&x))
	fieldOffset := uintptr(unsafe.Pointer(&dummy.r)) - uintptr(unsafe.Pointer(&dummy))
	return (*T)(unsafe.Pointer((uintptr)(unsafe.Pointer(t)) - fieldOffset))
}

func main() {
	//f := F{1, 1}
	//t := T{1, 1, []rune{'a'}}
	//s := t2(&t.r)
	//fmt.Println(s.a)
	//fmt.Println(t.a)

	//dataStructure.Test()

	//a := 1
	var a *int
	a = nil
	fmt.Println(a)

	m := make(map[int]int, 0)
}
