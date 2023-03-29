package main

import (
	"fmt"
)

func main() {
	name := "aa"

	switch name {
	case "Bima":
		fmt.Println("Hello bima")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Kenalan donk")
		fmt.Println("Maniss")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}
}
