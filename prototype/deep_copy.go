package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		a.StreetAddress,
		a.City,
		a.Country,
	}
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	q := *p
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}

// is a workable solution but is not the best
func DeepCopy() {
	john := Person{
		Name: "John",
		Address: &Address{
			StreetAddress: "123 London Rd",
			City:          "London",
			Country:       "UK",
		},
		Friends: []string{"Chris", "Matt"},
	}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"
	jane.Friends = append(jane.Friends, "Angela")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)

	// john := Person2{"John",
	// 	&Address1{
	// 		StreetAddress: "123 London Rd",
	// 		City:          "London",
	// 		Country:       "UK",
	// 	},
	// 	[]string{"Chris", "Matt"},
	// }
	//
	// // jane := john
	// // jane.Name = "Jane"                          // Ok, no problem here
	// // jane.Address1.StreetAddress = "321 Baker St" // Here is a problem, because Address1 wasn't be copied, they are sharing the same memory address for Address1 struct
	//
	// // deep copying
	//
	// // a way to solve this problem, not the best
	// jane := john
	// jane.Address1 = john.Address1.DeepCopy()
	//
	// // to this point both have the same address
	// fmt.Println(john, john.Address1)
	// fmt.Println(jane, jane.Address1)
}
