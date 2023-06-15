package dataStructure

import (
	"strings"
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

// SdsLen 已使用的长度
func SdsLen(s sds) int {
	sh := getSdshdr(s)
	return sh.len
}

// SdsAvail 未使用的长度
func SdsAvail(s sds) int {
	sh := getSdshdr(s)
	return sh.free
}

// SdsDup 复制一个sds
func SdsDup(s sds) sds {
	ts := string(*s)
	return SdsNewLen(ts, SdsLen(s))
}

// SdsRange 保存sds给定区间内的内容
// start和end都是闭区间
//start和end可以是负数，-1代表最后一个字符，-2代表倒数第二个自读，依此类推
func SdsRange(s sds, start int, end int) {
	sh := getSdshdr(s)
	newLen, len := 0, SdsLen(s)
	if len == 0 {
		return
	}
	if start < 0 {
		start += len
		if start < 0 {
			start = 0
		}
	}
	if end < 0 {
		end = len + end
		if end < 0 {
			end = 0
		}
	}
	if start > end {
		newLen = 0
	} else {
		newLen = end - start + 1
	}
	//如果start太大了不合法，就不能截取
	//如果len太大了不合法，就截取到最后一位
	//
	if newLen != 0 {
		if start >= len {
			newLen = 0
		} else if end >= len {
			end = len - 1
			if start > end {
				newLen = 0
			} else {
				newLen = end - start + 1
			}
		}
	} else {
		start = 0
	}
	if start != 0 && newLen != 0 {
		sh.buf = sh.buf[start : end+1]
		sh.free = sh.free + (sh.len - newLen)
		sh.len = newLen
	} else {
		sh.buf = []rune{}
		sh.free = sh.len
		sh.len = newLen
	}
}

//删除sds左右两端在cSet中出现的字符
func SdsTrim(s sds, cSet string) {
	sh := getSdshdr(s)
	newLen := 0
	i, start, j, end := 0, 0, SdsLen(s)-1, SdsLen(s)-1
	for i < end && strings.Contains(cSet, string((*s)[i])) {
		i++
	}
	for j >= start && strings.Contains(cSet, string((*s)[j])) {
		j--
	}
	if i > j {
		newLen = 0
	} else {
		newLen = j - i + 1
	}
	if newLen > 0 {
		sh.buf = sh.buf[i : j+1]
		sh.free = sh.free + sh.len - newLen
		sh.len = newLen
	}
}
