/*
 * File: payroll.go
 * Author: Franz Zbinden
 * Date: 09/30/2025
 * Purpose: Part 3 of the Assigment.
 */

package main

import "fmt"

// Employee basic info Struct
type employeeInfo struct {
	id, name, dpt, typeEmp string
}

// Employee type hourly
type employeeH struct {
	hourlyPayRate float64
	hours         int
	employeeInfo
}

// Employee type Sales
type employeeS struct {
	commission, sales float64
	employeeInfo
}

// Calculates weekly salary of hourly employee
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

// Calculates weekly salary of Sales employee
func (emp employeeS) weeklySalary() float64 {
	return emp.sales * emp.commission
}

// Employee Interface
type employee interface {
	weeklySalary() float64
}

// Printing data funcion for Hourly/Sales
func printData(emp employee) {
	switch e := emp.(type) {
	case employeeS:
		fmt.Printf("Type: %s\n"+
			"Id: %s\n"+
			"Name: %s\n"+
			"Department: %s\n"+
			"Commission: %.2f\n"+
			"Sales: $%.2f\n"+
			"Weekly Salary: $%.2f\n\n", e.typeEmp, e.id, e.name, e.dpt, e.commission, e.sales, e.weeklySalary())
	case employeeH:
		fmt.Printf("Type: %s\n"+
			"Id: %s\n"+
			"Name: %s\n"+
			"Department: %s\n"+
			"Pay Rate: $%.2f\n"+
			"Hours Worked: %d\n"+
			"Weekly Salary: $%.2f\n\n", e.typeEmp, e.id, e.name, e.dpt, e.hourlyPayRate, e.hours, e.weeklySalary())

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
		employeeS{
			employeeInfo: employeeInfo{id: "2222", name: "Jane Doe", dpt: "Sales", typeEmp: "Sales"},
			commission:   0.15,
			sales:        5000.00,
		},
		employeeH{
			employeeInfo:  employeeInfo{id: "3333", name: "Carlos Perez", dpt: "Software Engineering", typeEmp: "Hourly"},
			hourlyPayRate: 15.00,
			hours:         45,
		},
		employeeS{
			employeeInfo: employeeInfo{id: "4444", name: "Pepe Perez", dpt: "Marketing", typeEmp: "Sales"},
			commission:   0.12,
			sales:        8000.00,
		},
		employeeH{
			employeeInfo:  employeeInfo{id: "5555", name: "Juan Carlos", dpt: "Human Resources", typeEmp: "Hourly"},
			hourlyPayRate: 18.50,
			hours:         38,
		},
	}
	fmt.Println("These are the employees:\n")
	for amount, emp := range employees {
		printData(emp)
		fmt.Printf("There are %d employees", amount+1)
	}
}
