package test

import (
	"godis/dataStructure"
	"reflect"
	"testing"
)

func TestSdsNew(t *testing.T) {
	sds := dataStructure.SdsNew("foo")
	r := []rune{'f', 'o', 'o'}

	if dataStructure.SdsLen(sds) == 3 && reflect.DeepEqual(*sds, r) {
		t.Log("create a string and obtain the length")
	}

	sds1 := dataStructure.SdsNewLen("foo", 2)
	r1 := []rune{'f', 'o'}
	if dataStructure.SdsLen(sds1) == 2 && reflect.DeepEqual(*sds1, r1) {
		t.Log("Create a string with specified length")
	}
}
