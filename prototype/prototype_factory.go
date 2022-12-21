package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

var (
	mainOffice = Employee{
		Office: Address2{0, "123 East Dr", "London"},
	}

	auxOffice = Employee{
		Office: Address2{0, "66 West Dr", "London"},
	}
)

type Address2 struct {
	Suite               int
	StreetAddress, City string
}

type Employee struct {
	Name   string
	Office Address2
}

func newEmployee(proto *Employee, name string, suite int) *Employee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite
	return result
}

func NewMainOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&mainOffice, name, suite)
}

func NewAuxOfficeEmployee(name string, suite int) *Employee {
	return newEmployee(&auxOffice, name, suite)
}

func (p *Employee) DeepCopy() *Employee {
	var b bytes.Buffer
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)
	var result Employee
	_ = d.Decode(&result)
	return &result
}

func RunPrototypeFactory() {
	john := NewMainOfficeEmployee("John", 100)
	jane := NewMainOfficeEmployee("jane", 100)

	fmt.Println(john)
	fmt.Println(jane)
}
