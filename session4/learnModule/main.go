package main

import (
	"fmt"
	"countSalary"
	"countPayroll"
)

func main() {
	totalSalary := countSalary.CountTotalSalary(40,18000)
	payrollSalary := countPayroll.CountPayrollSalary(1,15000)

	salary := totalSalary-payrollSalary

	fmt.Println("total     : IDR ",totalSalary)
	fmt.Println("payroll   : IDR ",payrollSalary)
	fmt.Println("salary    : IDR ",salary)
}