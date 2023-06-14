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

	sds = dataStructure.SdsNewLen("foo", 2)
	r = []rune{'f', 'o'}
	if dataStructure.SdsLen(sds) == 2 && reflect.DeepEqual(*sds, r) {
		t.Log("Create a string with specified length")
	}
}
