package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

//while using embedded struct we don't have use the field name if we don't want to, go will assume the field name is same as the field type(contactInfo in this scenario)
type person struct {
	firstName string
	lastName  string
	contactInfo
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
		contactInfo: contactInfo{
			email:   "jim@ggmail.com",
			zipCode: 90000,
		},
	}

	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

//Pointters in go is relatively straight forward than it is on other programming languages like c, c++ etc.
//Go is 'Passed by value' language. This means that whenever we pass some value into a function.
//Go will take that value/struct, its going to copy all of that data inside that field
//and then place it inside a new container inside of our computers memory
//to be able to update our struct the way we want to update, we have to use pointers
//&variable ---> give me the memory adress of the value this variable is pointing at.
//*pointer ---> give me the value this memory adress is pointing at
//(pointerToPerson *person) --->in the reciever this is a type description, it means we are working with a pointer to a person
//*pointerToPerson ---> this is an operator - it means we want to manipulate the value the pointer is referencing

//when we thing about our variables here everthing thats going on with that struct that we want to update,
//we are really working with two different types of variables: pointers(adresses in memory) and value
//when we start thinigng about pointers and how they relate to the values that they contained within; here is the rule to really remember
//Turn pointers into value with *pointer
//Turn value into addresses with &value

//short cut for using receiver type of a pointer
//Go will automaticly gets the pointer when you call it like that jim.updateName("Jimmy")
//
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

//Gotchas with pointers - reference vs value types
//Slice vs arrays
//when we create a slice, go will automatically create an array and a structure that records the lenght of the slice, the capacity of the slice, and a reference to the underlying array
//in slice data structure there is a pointer to actual underliying array(pointer to array -  that keeps the values), capacity and length
//when you send a slice to a function, it copies the slice data structure to use in that function. But even though that data structure is copied, pointer to the actual value is still the same

// func main() {
// 	mySlice := []string{"Hi", "there", "how", "are", "you"}
// 	updateSlice(mySlice)
// 	fmt.Println(mySlice)

// }

// func updateSlice(s []string) {
// 	s[0] = "Bye"
// }

//reference types ---> slices, maps, channels, pointers, functions
//don't worry about pointers with reference types

//value types ---> int, float, string, bool, struct
//use pointers to change these value types in functions
