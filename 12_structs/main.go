package main

import ("fmt"
"strconv")

//Define person struct

type Person struct {
	// firstName string
	// lastName string
	// city string
	// gender string
	// age int

	firstName, lastName, city, gender string
	age int
}

//Greeting method (value reciever) nothing changed
func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

//hasBirthday method (pointer reciever) change is done
func (p *Person) hasBirthday() {
	p.age++
}

//getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	//Init person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 30}

	//Alternative
	person2 := Person{"Bob", "Johnson", "Seattle", "m", 55}

	fmt.Println(person1)
	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1.age)

	person1.hasBirthday()
	person1.getMarried("Williams")
	person2.getMarried("Thompson")

	fmt.Println(person2.greet())
}