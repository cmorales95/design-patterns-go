package main

import "fmt"

type Employee struct {
	Name, Position string
	AnnualIncome   int
}

// functional
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

type EmployeeFactory struct {
	Position     string
	AnnualIncome int
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnnualIncome}
}

func FactoryGenerator() {
	// NewEmployee(1)
	// this option is more recommended by the author of the class
	developerFactory := NewEmployeeFactory("Developer", 600000)
	managerFactory := NewEmployeeFactory("Manager", 80000)

	developer1 := developerFactory("Adam")
	manager1 := managerFactory("Jane")

	fmt.Println(developer1)
	fmt.Println(manager1)

	bossFactory := NewEmployeeFactory2("CEO", 100000)
	bossFactory.AnnualIncome = 110000 // real advantage is you could modify the values of the factory, in the old case you couldn't
	boss1 := bossFactory.Create("Sam")
	fmt.Println(boss1)
}
