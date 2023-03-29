package main

import "fmt"

func main() {
	name := "Bimammama"

	if name == "Bima" {
		fmt.Println("Hello bima")
	} else if name == "Joko" {
		fmt.Println("Hello bima")
	} else {
		fmt.Println("Hi, boleh kenalan?")
	}

	// short statement
	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

}
