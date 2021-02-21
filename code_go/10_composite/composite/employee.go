package composite

type GetEmployeeInfo interface {
	GetName() string
	GetJobNumber() int
}

type Employee struct {
	Name string
	JobNumber int
}

func NewEmployee(name string, num int) *Employee {
	return &Employee{
		Name: name,
		JobNumber: num,
	}
}

func (e *Employee) GetName() string {
	return e.Name
}

func (e *Employee) GetJobNumber() int {
	return e.JobNumber
}