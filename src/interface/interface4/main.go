package main

// 1. interface is a type, with 0 or more methods defined.
// 2. types that implement the interface can then be assign to  the interface variable

import "fmt"

type salaryCalculator interface {
	CalculateSalary() int
}

type Contract struct {
	empId    int
	basicpay int
}

type Permanent struct {
	empId    int
	basicpay int
	jj       int
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.jj
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}

func totalExpense(s []salaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total expense %d ", expense)
}

func main() {
	pemp1 := Permanent{1, 3000, 10000}
	pemp2 := Permanent{2, 3000, 20000}
	cemp1 := Contract{3, 3000}
	employees := []salaryCalculator{pemp1, pemp2, cemp1}
	totalExpense(employees)
}
