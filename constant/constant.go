package main

import "fmt"

func main() {
	const name string = "Bima"
	const friendName = "Pratandra"

	fmt.Println(name)
	fmt.Println(friendName)

	const (
		firstName string = "Rafli"
		lastName         = "Bima"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
