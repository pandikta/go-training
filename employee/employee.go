package employee

type Employee struct {
	EmpNo int
	Name  string
}

func (e Employee) getName() string {
	return e.Name
}
