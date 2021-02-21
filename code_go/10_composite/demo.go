package main

import (
	"fmt"
	"github.com/ZouXinyao/design_pattern/code_go/10_composite/composite"
)

func main() {
	company := composite.NewCompany("Company00")

	department01 := composite.NewDepartment("department01")
	department02 := composite.NewDepartment("department02")

	employee01 := composite.NewEmployee("Name01", 01)
	employee02 := composite.NewEmployee("Name02", 02)
	employee03 := composite.NewEmployee("Name03", 03)
	employee04 := composite.NewEmployee("Name04", 04)
	employee05 := composite.NewEmployee("Name05", 05)

	department01.Add(employee01)
	department01.Add(employee02)
	department02.Add(employee03)
	department02.Add(employee04)
	department02.Add(employee05)

	company.Add(department01)
	company.Add(department02)

	fmt.Println(company.SearchJobNumber(04))
	fmt.Println(company.SearchJobNumber(06))

	fmt.Println(company.SearchName("Name04"))
	fmt.Println(company.SearchName("Name00"))
}
