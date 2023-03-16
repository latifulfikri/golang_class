package main

import (
	"fmt"
	"net/http"
	"countSalary"
	"countPayroll"
)

func main()  {
	mysalary := salary()

	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "welcome to my server")
	})

	http.HandleFunc("/about",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hi! I am Fikri")
	})

	http.HandleFunc("/salary",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Your salary : %d", mysalary)
	})

	http.ListenAndServe("localhost:5050",nil)
}

func salary() int {
	totalSalary := countSalary.CountTotalSalary(40,18000)
	payrollSalary := countPayroll.CountPayrollSalary(1,15000)
	
	salary := totalSalary-payrollSalary

	return salary
}