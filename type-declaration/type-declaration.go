package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpBima NoKTP = "2183183718371"
	var marriedStatus = false

	fmt.Println(noKtpBima)
	fmt.Println(marriedStatus)
}
