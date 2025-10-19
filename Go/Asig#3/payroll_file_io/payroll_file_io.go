/*
 * File: payroll_file_io.go
 * Author: Franz Zbinden García
 * Course: COTI 4039-LH1
 * Date: 10/18/2025
 * Purpose: This Program reads a file of five employees of different types and creates a payroll
 * file with their data and weekly salary.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type employee struct {
	id         int
	name       string
	department string
	typee      string
}

type hourly struct {
	rate     float64
	hours    int
	emp_info employee
}

type sales struct {
	commission   float64
	sales_amount float64
	emp_info     employee
}

func parse_employee_line_hourly(line string) hourly {
	fields := strings.Split(line, "|")
	emp_type := strings.TrimSpace(fields[0])
	id, err := strconv.Atoi(strings.TrimSpace(fields[1]))
	if err != nil {
		panic("Error converting to int")
	}
	name := strings.TrimSpace(fields[2])
	dept := strings.TrimSpace(fields[3])

	rate_info, err := strconv.ParseFloat(strings.TrimSpace(fields[4]), 64)
	if err != nil {
		panic("Error converting salary_a")
	}
	hours_info, err := strconv.ParseFloat(strings.TrimSpace(fields[5]), 64)
	if err != nil {
		panic("Error converting salary_b")
	}

	return hourly{rate: rate_info, hours: int(hours_info), emp_info: employee{id: id, name: name, department: dept, typee: emp_type}}
}

func parse_employee_line_sales(line string) sales {
	fields := strings.Split(line, "|")
	emp_type := strings.TrimSpace(fields[0])
	id, err := strconv.Atoi(strings.TrimSpace(fields[1]))
	if err != nil {
		panic("Error converting to int")
	}
	name := strings.TrimSpace(fields[2])
	dept := strings.TrimSpace(fields[3])

	commission_info, err := strconv.ParseFloat(strings.TrimSpace(fields[4]), 64)
	if err != nil {
		panic("Error converting salary_a")
	}
	sales_amm_data, err := strconv.ParseFloat(strings.TrimSpace(fields[5]), 64)
	if err != nil {
		panic("Error converting salary_b")
	}

	return sales{commission: commission_info, sales_amount: sales_amm_data, emp_info: employee{id: id, name: name, department: dept, typee: emp_type}}
}

func (h hourly) calc_salary() float64 {
	return h.rate * float64(h.hours)
}

func (s sales) calc_salary() float64 {
	return s.commission * s.sales_amount
}

type employeess interface {
	calc_salary() float64
	new_employee_line() string
}

func employeetype(line string) string {
	fields := strings.Split(line, "|")
	return fields[0]
}

func (emp hourly) new_employee_line() string {
	return fmt.Sprintf("%s|%d|%s|%s|%.2f", emp.emp_info.typee, emp.emp_info.id, emp.emp_info.name, emp.emp_info.department, emp.calc_salary())
}

func (emp sales) new_employee_line() string {
	return fmt.Sprintf("%s|%d|%s|%s|%.2f", emp.emp_info.typee, emp.emp_info.id, emp.emp_info.name, emp.emp_info.department, emp.calc_salary())
}

func open_file_in(in_file_name string) *os.File {
	in_file, err := os.Open(in_file_name)
	if err != nil {
		panic("Error reading input file: " + in_file_name)
	}
	return in_file
}

func write_file_out(out_file_name string) *os.File {
	out_file, err := os.Create(out_file_name)
	if err != nil {
		panic("Error closing file: " + out_file_name)
	}
	return out_file
}

func main() {
	in_file := open_file_in("employees.txt")
	out_file := write_file_out("payroll.txt")

	defer in_file.Close()
	defer out_file.Close()

	// Create scanner
	scanner := bufio.NewScanner(in_file)

	for scanner.Scan() {
		line := scanner.Text()
		employee_type := employeetype(line)

		if employee_type == "H" {
			employee := parse_employee_line_hourly(line)
			fmt.Fprintln(out_file, employee.new_employee_line())
		} else if employee_type == "S" {
			employee := parse_employee_line_sales(line)
			fmt.Fprintln(out_file, employee.new_employee_line())
		}
	}
}
