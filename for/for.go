package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	slice := []string{"Rafli", "Bima", "Pratandra"}

	for i, v := range slice {
		fmt.Println("Index", i, "=", v)
	}

	person := make(map[string]string)

	person["name"] = "Rafli"
	person["title"] = "Programmer"

	for key, v := range person {
		fmt.Println(key, "=", v)
	}

}
