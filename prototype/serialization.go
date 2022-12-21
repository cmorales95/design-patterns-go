package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Address1 struct {
	StreetAddress, City, Country string
}

type Person1 struct {
	Name    string
	Address *Address1
	Friends []string
}

func (p *Person1) DeepCopy() *Person1 {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	_ = e.Encode(p)

	fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)
	result := Person1{}
	_ = d.Decode(&result)
	return &result
}

// is a workable solution but is not the best
func DeepCopy1() {
	john := Person1{
		Name: "John",
		Address: &Address1{
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
	// jane.Address1 = john.Address1.DeepCopy1()
	//
	// // to this point both have the same address
	// fmt.Println(john, john.Address1)
	// fmt.Println(jane, jane.Address1)
}
