package composite

import "strings"

type SearchEmployee interface {
	SearchName(name string) bool
	SearchJobNumber(num int) bool
}

type DepartmentSkill interface {
	Add(e Employee)
	GetDepartmentName() string
}

type Department struct {
	Name string
	Employees []*Employee
}

func NewDepartment(name string) *Department {
	return &Department{
		Name: name,
	}
}

func (d *Department) GetDepartmentName() string {
	return d.Name
}

func (d *Department) Add(e *Employee) {
	d.Employees = append(d.Employees, e)
}

func (d *Department) SearchName(name string) bool {
	for _, employee := range d.Employees {
		if strings.EqualFold(name, employee.GetName()) {
			return true
		}
	}
	return false
}

func (d *Department) SearchJobNumber(num int) bool {
	for _, employee := range d.Employees {
		if  num == employee.GetJobNumber() {
			return true
		}
	}
	return false
}