package solid

// Dependency Inversion Principle
// High Level Modules should not depend on Low Level Modules
// Both should depend on abstractions
import "fmt"

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	name string
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

// low-level module
type Relationships struct {
	relations []Info
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)
	for i, v := range r.relations {
		if v.relationship == Parent && v.from.name == name {
			result = append(result, r.relations[i].to)
		}
	}

	return result
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relations = append(r.relations, Info{
		from:         parent,
		relationship: Parent,
		to:           child,
	})

	r.relations = append(r.relations, Info{
		from:         child,
		relationship: Child,
		to:           parent,
	})
}

// high-lovel-module
type Research struct {
	// break DIP
	// relationships Relationships
	browser RelationshipBrowser
}

func (r *Research) Investigate() {
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called", p.name)
	}
}

func ExecuteDIP() {
	parent := Person{name: "John"}
	child1 := Person{name: "Chris"}
	child2 := Person{name: "Matt"}

	var relationships Relationships
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	r := Research{&relationships}
	r.Investigate()

}
