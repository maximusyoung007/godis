package dataStructure

//用于指向sdshdr的buf属性
type sds *rune

type sdshdr struct {
	len  int
	free int
	buf  []rune
}

func sdsnew(s string) sdshdr {

}
