package composite

type CompanySkill interface {
	Add(e Employee)
	GetCompanyName() string
}

type Company struct {
	Name string
	Department []*Department
}

func NewCompany(name string) *Company {
	return &Company{
		Name: name,
	}
}

func (c *Company) GetCompanyName() string {
	return c.Name
}

func (c *Company) Add(d *Department) {
	c.Department = append(c.Department, d)
}

func (c *Company) SearchName(name string) bool {
	for _, depart := range c.Department {
		if depart.SearchName(name) {
			return true
		}
	}
	return false
}

func (c *Company) SearchJobNumber(num int) bool {
	for _, depart := range c.Department {
		if depart.SearchJobNumber(num) {
			return true
		}
	}
	return false
}
