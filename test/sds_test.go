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

func TestSdsTrim(t *testing.T) {
	s := dataStructure.SdsNew("AA...AA.a.aa.aHelloWorld     :::")
	dataStructure.SdsTrim(s, "Aa. :")
	if reflect.DeepEqual(string(*s), "HelloWorld") {
		t.Log("test sdsTrim1")
	} else {
		t.Fatalf("err: %s", string(*s))
	}
}

func TestSdsRange(t *testing.T) {
	s := dataStructure.SdsNew("Hello World")
	dataStructure.SdsRange(s, 1, -1)
	if reflect.DeepEqual(*s, []rune("ello World")) {
		t.Log("test sdsRange 1")
	}

	x := dataStructure.SdsNew("xxciaoyyy")
	dataStructure.SdsTrim(x, "xy")
	if reflect.DeepEqual(string(*x), "ciao") && dataStructure.SdsLen(x) == 4 {
		t.Log("sdstrim() correctly trims characters")
	} else {
		t.Fatalf("err: %s", string(*x))
		t.Fatalf("err: %d", dataStructure.SdsLen(x))
	}

	y := dataStructure.SdsDup(x)
	dataStructure.SdsRange(y, 1, 1)
	if reflect.DeepEqual(string(*y), "i") && dataStructure.SdsLen(y) == 1 {
		t.Log("sdsrange(...,1,1)")
	} else {
		t.Fatalf("err: %s ----- %d", string(*y), dataStructure.SdsLen(y))
	}

	y = dataStructure.SdsDup(x)
	dataStructure.SdsRange(y, 1, -1)
	if reflect.DeepEqual(string(*y), "iao") && dataStructure.SdsLen(y) == 3 {
		t.Log("sdsrange(...,1,-1)")
	} else {
		t.Fatalf("err: %s ----- %d", string(*y), dataStructure.SdsLen(y))
	}

	y = dataStructure.SdsDup(x)
	dataStructure.SdsRange(y, -2, -1)
	if reflect.DeepEqual(string(*y), "ao") && dataStructure.SdsLen(y) == 2 {
		t.Log("sdsrange(...,-2,-1)")
	} else {
		t.Fatalf("err: %s ----- %d", string(*y), dataStructure.SdsLen(y))
	}

	y = dataStructure.SdsDup(x)
	dataStructure.SdsRange(y, 2, 1)
	if reflect.DeepEqual(string(*y), "") && dataStructure.SdsLen(y) == 0 {
		t.Log("sdsrange(...,2,1)")
	} else {
		t.Fatalf("err: %s ----- %d", string(*y), dataStructure.SdsLen(y))
	}

	y = dataStructure.SdsDup(x)
	dataStructure.SdsRange(y, 1, 100)
	if reflect.DeepEqual(string(*y), "iao") && dataStructure.SdsLen(y) == 3 {
		t.Log("sdsrange(...,1,100)")
	} else {
		t.Fatalf("err: %s ----- %d", string(*y), dataStructure.SdsLen(y))
	}

	y = dataStructure.SdsDup(x)
	dataStructure.SdsRange(y, 100, 100)
	if reflect.DeepEqual(string(*y), "") && dataStructure.SdsLen(y) == 0 {
		t.Log("sdsrange(...,100,100)")
	} else {
		t.Fatalf("err: %s ----- %d", string(*y), dataStructure.SdsLen(y))
	}
}

func TestSdsMakeRoomFor(t *testing.T) {
	x := dataStructure.SdsNew("0")
	x = dataStructure.SdsMakeRoomFor(x, 1)
	if dataStructure.SdsLen(x) == 1 && dataStructure.SdsAvail(x) > 0 {
		t.Log("sdsMakeRoomFor()")
	} else {
		t.Fatalf("err: #{x}")
	}
}

func TestSdsCat(t *testing.T) {
	x := dataStructure.SdsNewLen("foo", 2)
	x = dataStructure.SdsCat(x, "bar")
	if reflect.DeepEqual(string(*x), "fobar") && dataStructure.SdsLen(x) == 5 {
		t.Log("Strings concatenation")
	} else {
		t.Fatalf("err: %s ----- %d", string(*x), dataStructure.SdsLen(x))
	}

	x = dataStructure.SdsCpy(x, "a")
	if reflect.DeepEqual(string(*x), "a") && dataStructure.SdsLen(x) == 1 {
		t.Log("sdscpy() against an originally longer string")
	} else {
		t.Fatalf("err: %s ----- %d", string(*x), dataStructure.SdsLen(x))
	}

	x = dataStructure.SdsCpy(x, "xyzxxxxxxxxxxyyyyyyyyyykkkkkkkkkk")
	if reflect.DeepEqual(string(*x), "xyzxxxxxxxxxxyyyyyyyyyykkkkkkkkkk") && dataStructure.SdsLen(x) == 33 {
		t.Log("sdscpy() against an originally shorter string")
	} else {
		t.Fatalf("err: %s ----- %d", string(*x), dataStructure.SdsLen(x))
	}
}

func TestSdsCatSds(t *testing.T) {
	x := dataStructure.SdsNewLen("foo", 2)
	y := dataStructure.SdsNewLen("bar", 2)
	z := dataStructure.SdsCatSds(x, y)
	if reflect.DeepEqual(string(*z), "foba") && dataStructure.SdsLen(z) == 4 {
		t.Log("sds concatenation")
	} else {
		t.Fatalf("err: %s ----- %d", string(*z), dataStructure.SdsLen(z))
	}
}

func TestSdsCmp(t *testing.T) {
	x := dataStructure.SdsNew("foo")
	y := dataStructure.SdsNew("foa")
	if dataStructure.SdsCmp(x, y) > 0 {
		t.Log("sdscmp(foo, foa), large")
	} else {
		t.Fatalf("not larger")
	}

	x = dataStructure.SdsNew("bar")
	y = dataStructure.SdsNew("bar")
	if dataStructure.SdsCmp(x, y) == 0 {
		t.Log("sdscmp(bar, bar), equal")
	} else {
		t.Fatalf("not equal")
	}

	x = dataStructure.SdsNew("aar")
	y = dataStructure.SdsNew("bar")
	if dataStructure.SdsCmp(x, y) < 0 {
		t.Log("sdscmp(aar, bar), small")
	} else {
		t.Fatalf("not smaller")
	}
}
