package dataStructure

import (
	"unsafe"
)

//用于指向sdshdr的buf属性
type sds *[]rune

type sdshdr struct {
	//长度
	len int
	//预留空间
	free int
	//字符数组
	buf []rune
}

// SdsNew 创建一个包含字符串s的sds
func SdsNew(s string) sds {
	initLen := len(s)
	return SdsNewLen(s, initLen)
}

// SdsEmpty 创建一个空的sds
func SdsEmpty() sds {
	return SdsNewLen("", 0)
}

func SdsNewLen(s string, initLen int) sds {
	r := make([]rune, 0)
	for i := 0; i < initLen; i++ {
		r = append(r, rune(s[i]))
	}
	sh := sdshdr{initLen, 0, r}
	return &sh.buf
}

//func sdsFree(s sds) {}

//根据指向sds的指针偏移到指向sdshdr的指针
func getSdshdr(s sds) *sdshdr {
	var dummy sdshdr
	//计算偏移量
	fieldOffSet := uintptr(unsafe.Pointer(&dummy.buf)) - uintptr(unsafe.Pointer(&dummy))
	return (*sdshdr)(unsafe.Pointer((uintptr)(unsafe.Pointer(s)) - fieldOffSet))
}

//已使用的长度
func SdsLen(s sds) int {
	sh := getSdshdr(s)
	return sh.len
}

//未使用的长度
func SdsAvail(s sds) int {
	sh := getSdshdr(s)
	return sh.free
}
