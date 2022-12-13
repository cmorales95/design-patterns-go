package main

import "fmt"

type Person interface {
	SayHello()
}

type person struct { // is lowercase
	name string
	age  int
}

type TiredPerson struct {
	name string
	age  int
}

func (p *person) SayHello() {
	fmt.Printf("Hi, my name is %s, I am %d years old\n", p.name, p.age)
}

func (p *TiredPerson) SayHello() {
	fmt.Printf("Sorry, I'm too tired")
}

func NewPerson1(name string, age int) Person {
	if age > 100 { // returning another struct
		return &TiredPerson{name, age}
	}
	return &person{name, age}
}

func ExecuteFactoryInterface() {
	p := NewPerson1("James", 134)
	p.SayHello()
}
