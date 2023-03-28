package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "bima",
		"address": "depok",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	book := make(map[string]string)

	book["title"] = "Belajar golang"
	book["author"] = "Pzn"
	book["ups"] = "salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}
