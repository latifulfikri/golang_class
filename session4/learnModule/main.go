package main

import (
	"fmt"
	"countSalary"
)

func main() {
	totalSalary := countSalary.CountTotalSalary(40,18000)
	fmt.Println(totalSalary)
}