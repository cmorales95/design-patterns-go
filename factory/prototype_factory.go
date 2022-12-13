package main

import "fmt"

type Employee1 struct {
	Name, Position string
	AnnualIncome   int
}

const (
	Developer = iota
	Manager
)

func NewEmployee(role int) *Employee1 {
	switch role {
	case Developer:
		return &Employee1{"", "developer", 60000}
	case Manager:
		return &Employee1{"", "manager", 80000}
	default:
		panic("unsupported role")
	}
}

func PrototypeFactory() {
	m := NewEmployee(Manager)
	m.Name = "Sam"
	fmt.Println(m)
}
