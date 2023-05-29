package employee

type Manager struct {
	// isManager bool
	Title string
}

type Programmer struct {
	Employee
	Manager
	CodeEditor string
}

func NewProgrammer(name string, CodeEditor string) Programmer {
	return Programmer{
		Employee: Employee{
			EmpNo: 1,
			Name:  name,
		},
		Manager: Manager{
			Title: "Manager Dev",
		},
		CodeEditor: CodeEditor,
	}
}
