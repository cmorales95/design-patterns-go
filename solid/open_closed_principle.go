package solid

import "fmt"

// combination of OCP and Repository demo

type Color int

const (
	red Color = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

func (f *Filter) filterByColor(
	products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

func (f *Filter) filterBySize(
	products []Product, size Size) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size {
			result = append(result, &products[i])
		}
	}

	return result
}

func (f *Filter) filterBySizeAndColor(
	products []Product, size Size,
	color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.size == size && v.color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

// Specification Pattern
type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == c.color
}

type SizeSpecification struct {
	size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == s.size
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, product := range products {
		if spec.IsSatisfied(&product) {
			result = append(result, &products[i])
		}
	}

	return result
}

// AndSpecification , Composite Design Pattern.. a little example
type AndSpecification struct {
	first, second Specification
}

func (j AndSpecification) IsSatisfied(p *Product) bool {
	return j.first.IsSatisfied(p) && j.second.IsSatisfied(p)
}

func ExecutesOCP() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	fmt.Printf("Green products (old):\n")
	var f Filter
	for _, v := range f.filterByColor(products, green) {
		fmt.Printf(" - %s is green \n", v.name)
	}
	/*
		OLD

		NEW
	*/

	fmt.Printf("Green products (new):\n")
	greenSpec := ColorSpecification{green}
	var bf BetterFilter
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf(" - %s is green \n", v.name)
	}

	largeSpec := SizeSpecification{large}
	lgSpec := AndSpecification{greenSpec, largeSpec}
	fmt.Printf("Large Green products (new):\n")
	for _, v := range bf.Filter(products, lgSpec) {
		fmt.Printf(" - %s is large and green \n", v.name)
	}
}
