package main

import "fmt"

func main() {
	dataTypes()
}

func dataTypes() {
	var first uint8 = 89
	var second int8 = -12

	fmt.Printf("Tipe data first : %T\n", first)
	fmt.Printf("Tipe data second : %T\n", second)

	var decimalNumber float32 = 3.14

	fmt.Printf("decimal Number: %f\n", decimalNumber)
	fmt.Printf("decimal Number: %.3f\n", decimalNumber)

	var kondisi bool = true
	fmt.Printf("Apakah kamu seorang QA Engineer? %t\n", kondisi)

	var kalimat string = "Siap!!!"
	fmt.Print(kalimat)

	var kalimatPanjang = `Jadi begini caranya pake simbol "backticks"
		yang di bawah cacing/tilde "Okay?"
		`
	fmt.Print(kalimatPanjang)
}
