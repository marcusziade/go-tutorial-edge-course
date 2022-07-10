package main

type Engineer struct {
	Name    string
	Project Project
}

type Priority string

const (
	Top       Priority = "Top priority"
	Secondary Priority = "Secondary priority"
	Tertiary  Priority = "Tertiary priority"
)

type Project struct {
	Name     string
	Priority Priority
}

func (e *Engineer) UpdateName(name string) {
	e.Name = name
}

func (e *Engineer) GetProjectPriority() Priority {
	return e.Project.Priority
}

func (e Engineer) Print() {
	println("Engineer info")
	println("Name:\n", e.Name)
	println("Project info\n")
	println(e.Project.Name, e.Project.Priority)
}

func main() {
	project := Project{
		Name:     "Project X",
		Priority: Top,
	}

	engineer := Engineer{
		Name:    "Marcus",
		Project: project,
	}

	println(engineer.GetProjectPriority())

	// println(engineer.Name)

	// engineer.UpdateName("Axl")
	// println(engineer.Name)

	// println(engineer.Name)
	// println(engineer.Project.Name)
	// println(engineer.Project.Priority)
	// fmt.Println(engineer)
	// println(engineer.Name)
	// More verbose
	// fmt.Printf("%+v\n", engineer)
}
