package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// var alex1 person
	// alex := person{"Alex", "Anderson"}
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
	alex.print()

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	jimpointer := &jim
	jimpointer.updateName("Jimmy")
	jim.print()
	jim.updateName("Jimme")
	jim.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
