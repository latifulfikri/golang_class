package main

import "fmt"

var a = 10

func calc(num int) int {
	num += 12
	return num
}

func main() {
	fmt.Println("Your angka is", calc(a))
}
