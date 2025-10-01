/*
 * File: payroll2.go
 * Author: Franz Zbinden Garcia
 * Course: COTI 4039-VH1
 * Date: 09/30/2025
 * Purpose: This program creates a slice of five employees of different types
 *          and displays their data and weekly salary.
 */

package main

import "fmt"

type employeeInfo struct {
	id, name, dpt, typeEmp string
}

type employeeH struct {
	hourlyPayRate float64
	hours         int
	employeeInfo
}

type employeeS struct {
	commission, sales float64
	employeeInfo
}

func (emp employeeH) weeklySalary() float64 {
	if emp.hours <= 0 {
		return 0
	}
	if emp.hours > 40 {
		overtimePayrate := emp.hourlyPayRate * 1.5
		overTimeSalary := (float64(emp.hours - 40)) * overtimePayrate
		normalSalary := float64(40) * float64(emp.hourlyPayRate)
		return overTimeSalary + normalSalary

	} else {
		return (float64(emp.hours) * float64(emp.hourlyPayRate))
	}
}

func (emp employeeS) weeklySalary() float64 {
	return emp.sales * emp.commission
}

type employee interface {
	weeklySalary() float64
}

func printData(emp employee) {
	switch e := emp.(type) {
	case employeeS:
		// fmt.Printf("%s (Hourly) | Rate: $%.2f | Hours: %d | Weekly Salary: $%.2f\n",
		//     e.name, e.hourlyPayRate, e.hours, e.weeklySalary())
	case employeeH:
		fmt.Printf("Type: %s\n"+
			"Id: %d\n"+
			"Name: %s\n"+
			"Department: %s\n"+
			"Pay Rate: %d\n"+
			"Hours Worked: %f\n"+
			"Weekly Salary: $%f", emp.typeEmp, emp.id, emp.name, emp.dpt, emp.hourlyPayRate, emp.hours)

	default:
		fmt.Println("Unknown employee type")
	}
}

func main() {
	employees := []employee{
		employeeH{
			employeeInfo:  employeeInfo{id: "1111", name: "John Doe", dpt: "Information Technology", typeEmp: "Hourly"},
			hourlyPayRate: 15.00,
			hours:         40,
		},
	}
	fmt.Println("These are the employees...")
	for _, emp := range employees {
		printData(emp)
	}
}
