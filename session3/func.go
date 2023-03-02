package main

import (
	"fmt"
)

func f() *int {
	v := 1
	fmt.Println(v)
	return &v
}

func main() {
	var p = f()
	fmt.Println(f())
	fmt.Println(p)
}
