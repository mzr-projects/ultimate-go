package composition

import "fmt"

type Employee struct {
	Name string
	ID   int
}

func (e Employee) description() string {
	return fmt.Sprintf("Employee, Name: %s, ID: %d", e.Name, e.ID)
}

/*
Manager contains Employee and no name is assigned to it, so that makes it an Embedded Field
*/
type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) description() string {
	return fmt.Sprintf("Manager: %s, Reports: %v", m.Employee.description(), m.Reports)
}

func SevenDemo() {
	m := Manager{
		Employee: Employee{
			Name: "John",
			ID:   1,
		},
		Reports: []Employee{
			{
				Name: "Jane",
				ID:   2,
			},
		},
	}

	fmt.Println(m.description())
}
