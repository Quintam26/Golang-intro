package main

import "fmt"

func main() {
	//Using var
	// var name = "Mario"
	var age = 28
	const isCool = true

	//Shorthand
	// name := "Mario"
	// email := "mquinta@gmail.com"

	name, email := "Mario", "mquinta@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", email)
}