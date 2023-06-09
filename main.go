package main

import "fmt"

func main() {
	fmt.Println("hello world")
	//s := "abc你好"
	var b []rune = []rune{'你', 's'}
	for _, v := range b {
		fmt.Println(v)
	}
	fmt.Println(len(b))
}
