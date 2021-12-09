package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

//all the different ways to create value of type struct

func main() {
	//go assumes the order of definition of fields
	//I personally can not stand this sytax, but you will see a lot of official code that only rely on the order of definition
	//relying upon the order of definition of fields
	// muge := person{"Müge", "Evren Bulut"} // if you ever swap the order of field, things will go nuts

	// muge := person{firstName: "Müge", lastName: "Evren Bulut"} //this is the safe way

	// var muge person
	// muge.firstName = "Müge"
	// muge.lastName = "Evren Bulut"
	// fmt.Println(muge)
	// fmt.Printf("%+v", muge)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@ggmail.com",
			zipCode: 90000,
		},
	}
	fmt.Printf("%+v", jim)

}
