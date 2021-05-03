package main

import "reflect"

type s struct {
	a string
	b int
}

func main() {
	s := []s{
		{"aa", 33},
		{"cc", 11},
		{"bb", 22},
	}
	v := reflect.ValueOf(s[0]).Elem().Type()
	println(v)
}
