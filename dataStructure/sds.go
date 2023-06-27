package dataStructure

import (
	"strings"
	"unsafe"
)

const SDS_MAX_PREALLOC int = 1024 * 1024

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
	r := make([]rune, initLen)
	for i := 0; i < initLen; i++ {
		r[i] = rune(s[i])
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

func SdsClear(s sds) {
	sh := getSdshdr(s)
	sh.free = sh.free + sh.len
	sh.len = 0
	sh.buf = []rune{}
}

func SdsCat(s sds, t string) sds {
	return SdsCatLen(s, t, len(t))
}

// SdsCatLen 将长度为len的字符串t追加到sds后
func SdsCatLen(s sds, t string, length int) sds {
	curLen := SdsLen(s)
	s = SdsMakeRoomFor(s, length)

	sh := getSdshdr(s)
	sh.len = curLen + length
	//右边的free为扩容以后的
	sh.free = sh.free - length
	k := 0
	for i := length - 1; i < sh.len && k < len(t); i++ {
		(*s)[i] = rune(t[k])
		k++
	}

	return &sh.buf
}

func SdsCatSds(s sds, t sds) sds {
	str := string(*t)
	return SdsCatLen(s, str, len(str))
}

// SdsMakeRoomFor 字符串扩容,拥有addLen的空余长度
func SdsMakeRoomFor(s sds, addLen int) sds {
	var sh *sdshdr
	var newSh *sdshdr

	free := SdsAvail(s)
	if free > addLen {
		return s
	}

	var length, newLen int

	length = SdsLen(s)
	sh = getSdshdr(s)

	newLen = length + addLen
	newSh = sh
	newSh.len = newLen

	if newLen < SDS_MAX_PREALLOC {
		newLen *= 2
	} else {
		newLen += SDS_MAX_PREALLOC
	}

	newS := make([]rune, newSh.len, newLen)
	for i := 0; i < len(newSh.buf); i++ {
		newS[i] = newSh.buf[i]
	}

	newSh.buf = newS
	newSh.free = newLen - length

	return &newSh.buf
}

// SdsCpy 复制t到sds
func SdsCpy(s sds, t string) sds {
	return SdsCpyLen(s, t, len(t))
}

// SdsCpyLen 将t的前len个字符复制到s中，
func SdsCpyLen(s sds, t string, length int) sds {
	sh := getSdshdr(s)
	totalLen := sh.free + sh.len
	if totalLen < length {
		s = SdsMakeRoomFor(s, length-sh.len)
		sh = getSdshdr(s)
		totalLen = sh.free + sh.len
	}

	sh.len = length
	sh.free = totalLen - length

	r := make([]rune, length, totalLen)
	for i := 0; i < length; i++ {
		r[i] = rune(t[i])
	}
	sh.buf = r

	return &sh.buf
}

// SdsGrowZero 用空字符串将sds扩展到指定长度
func SdsGrowZero(s sds, len int) sds {
	sh := getSdshdr(s)
	var curLen int
	if len < curLen {
		return s
	}
	s = SdsMakeRoomFor(s, len-curLen)
	sh = getSdshdr(s)
	sh.len = len
	return s
}

// SdsCmp 比较两个sds大小
func SdsCmp(s1 sds, s2 sds) int {
	l1 := SdsLen(s1)
	l2 := SdsLen(s2)
	i := 0
	for i < l1 && i < l2 {
		if (*s1)[i] > (*s2)[i] {
			return 1
		} else if (*s1)[i] < (*s2)[i] {
			return -1
		}
		i++
	}
	if l1-i > 0 {
		return 1
	}
	if l2-i > 0 {
		return -1
	}
	return 0
}
