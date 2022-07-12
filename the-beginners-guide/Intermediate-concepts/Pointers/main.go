package main

type Planet struct {
	Name string
	Age  int
}

func (p *Planet) UpdateName(name string) {
	p.Name = name
}

func main() {
	planet := Planet{Name: "Venus", Age: 4_000_000_000}
	println(planet.Name)
	planet.UpdateName("Saturn")
	println(planet.Name)
}
