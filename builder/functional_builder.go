package main

import "fmt"

type Person1 struct {
	Name, Postion string
}

type personMod func(person *Person1)
type PersonBuilder1 struct {
	actions []personMod // allow to extend the builder with more methods
}

func (b *PersonBuilder1) Called(name string) *PersonBuilder1 {
	b.actions = append(b.actions, func(person *Person1) {
		person.Name = name
	})

	return b
}

func (b *PersonBuilder1) Build() *Person1 {
	p := Person1{}
	for _, a := range b.actions {
		a(&p)
	}

	return &p
}

// Extending builder
func (b *PersonBuilder1) WorkAsA(position string) *PersonBuilder1 {
	b.actions = append(b.actions, func(person *Person1) {
		person.Postion = position
	})

	return b
}

func RunFunctionalBuilder() {
	b := PersonBuilder1{}
	p := b.Called("Dmitri").
		WorkAsA("Developer").
		Build()
	fmt.Println(*p)
}
