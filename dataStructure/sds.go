package dataStructure

import (
	"fmt"
	"unsafe"
)

//用于指向sdshdr的buf属性
type sds *rune

type sdshdr struct {
	//长度
	len int
	//预留空间
	free int
	//字符数组
	buf []rune
}

//创建一个包含字符串s的sds
func sdsNew(s string) sds {
	initLen := len(s)
	return sdsNewLen(s, initLen)
}

//创建一个空的sds
func sdsEmpty() sds {
	return sdsNewLen("", 0)
}

func sdsNewLen(s string, initLen int) sds {
	sh := &sdshdr{initLen, 0, []rune(s)}
	//将数组sh.buf的地址赋给指针sds
	fmt.Println(&sh)
	fmt.Println(&sh.buf[0])
	fmt.Println(&sh.buf[0] - &sh)
	return &sh.buf[0]
}

//func sdsFree(s sds) {}

//返回已使用的长度
func sdsLen(s sds) int {
	var sds1 sdshdr
	t := unsafe.Sizeof(sds1)
	var sds2 sdshdr
	t2 := unsafe.Sizeof(sds2)
	fmt.Println(t)
	fmt.Println(t2)
	unSafeS := unsafe.Pointer(s)
	unSafeSh := uintptr(unSafeS) - t
	sh := (*sdshdr)(unsafe.Pointer(unSafeSh))
	return sh.len
}

func Test() {
	sd := sdsNew("abc")
	fmt.Println(sdsLen(sd))
}
