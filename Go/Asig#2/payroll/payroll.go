/*
 * File: payroll.go
 * Author: Franz Zbinden Garcia
 * Course: COTI 4039-VH1
 * Purpose: This program creates a slice of five employees of different types
 *          and displays their data and weekly salary.
 */

package main

import "fmt"

// Represents the set of available departments.
type department int

const (
	finance department = iota
	humanResources
	informationTechnology
	marketing
	sales
)

// Returns the string representation of the department.
func (d department) String() string {
	if finance > d || d > sales {
		return "Unknown"
	}
	deptName := map[department]string{
		finance:               "Finance",
		humanResources:        "Human Resources",
		informationTechnology: "Information Technology",
		marketing:             "Marketing",
		sales:                 "Sales",
	}
	return deptName[d]
}

// Represents the common information for all employees.
type employeeInfo struct {
	id         string
	name       string
	department department
}

// Returns the string representation of the employee info.
func (ei employeeInfo) String() string {
	return fmt.Sprintf("ID: %s, Name: %s, Department: %s",
		ei.id, ei.name, ei.department)
}

// Represents the set of methods common to any employee.
type employee interface {
	weeklySalary() float64
}

// Represents an hourly employee with pay rate and hours worked.
type hourlyEmployee struct {
	employeeInfo
	hourlyPayRate float64
	hoursWorked   int
}

// Computes the weekly salary of the hourly employee.
func (he hourlyEmployee) weeklySalary() float64 {
	if he.hoursWorked <= 40 {
		return he.hourlyPayRate * float64(he.hoursWorked)
	}
	regularPay := he.hourlyPayRate * 40.0
	overtimeHours := float64(he.hoursWorked - 40)
	overtimePay := he.hourlyPayRate * 1.5 * overtimeHours
	return regularPay + overtimePay
}

// Returns the string representation of the hourly employee.
func (he hourlyEmployee) String() string {
	return fmt.Sprintf("%s, Hourly Pay Rate: $%.2f, Hours Worked: %d",
		he.employeeInfo, he.hourlyPayRate, he.hoursWorked)
}

// Represents a salesperson with commission rate and sales amount.
type salesperson struct {
	employeeInfo
	commissionRate float64
	salesAmount    float64
}

// Computes the weekly salary of the salesperson.
func (sp salesperson) weeklySalary() float64 {
	return sp.commissionRate * sp.salesAmount
}

// Returns the string representation of the salesperson.
func (sp salesperson) String() string {
	return fmt.Sprintf("%s, Commission Rate: %.2f%%, Sales Amount: $%.2f",
		sp.employeeInfo, sp.commissionRate*100.0, sp.salesAmount)
}

// Prints the data of the given employee.
func printData(emp employee) {
	fmt.Println("\n" + fmt.Sprint(emp))
	fmt.Printf("Weekly Salary: $%.2f\n", emp.weeklySalary())
}

// Starts the execution of the program.
func main() {
	employees := []employee{
		hourlyEmployee{
			employeeInfo:  employeeInfo{id: "1111", name: "John Doe", department: informationTechnology},
			hourlyPayRate: 15.00,
			hoursWorked:   40,
		},
		salesperson{
			employeeInfo:   employeeInfo{id: "2222", name: "Jane Doe", department: sales},
			commissionRate: 0.15,
			salesAmount:    5000.00,
		},
		hourlyEmployee{
			employeeInfo:  employeeInfo{id: "3333", name: "Bob Smith", department: finance},
			hourlyPayRate: 20.00,
			hoursWorked:   45,
		},
		salesperson{
			employeeInfo:   employeeInfo{id: "4444", name: "Alice Johnson", department: marketing},
			commissionRate: 0.12,
			salesAmount:    8000.00,
		},
		hourlyEmployee{
			employeeInfo:  employeeInfo{id: "5555", name: "Charlie Brown", department: humanResources},
			hourlyPayRate: 18.50,
			hoursWorked:   38,
		},
	}

	fmt.Println("These are the employees...")
	for _, emp := range employees {
		printData(emp)
	}
}
