package main

type Employee interface {
	GetName() string
}

type Engineer struct {
	Name string
}

type Manager struct {
	Name string
}

func (e *Engineer) GetName() string {
	return "Employee name: " + e.Name
}

func (m *Manager) GetName() string {
	return "Manager name: " + m.Name
}

func PrintDetails(e Employee) {
	println(e.GetName())
}

func main() {
	engineer := &Engineer{Name: "Marcus"}
	manager := &Manager{Name: "Voldemort"}
	PrintDetails(engineer)
	PrintDetails(manager)
}