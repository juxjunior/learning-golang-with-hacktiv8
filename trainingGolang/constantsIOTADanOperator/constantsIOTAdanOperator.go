package main

import "fmt"

func main() {
	constantsIOTAdanOperator()
}

func constantsIOTAdanOperator() {
	//Konstanta
	const phi = 3.14
	const notFoundCode = 404
	const applicationName = "Golang"
	const secretKey = "kjhfaw98yaw98y9awhf9hawf"

	fmt.Println(phi)
	fmt.Println(notFoundCode)
	fmt.Println(applicationName)
	fmt.Println(secretKey)

	//IOTA
	const (
		c1 = 404 + iota
		c2
		c3
		c4
	)

	fmt.Println(c1, c2, c3, c4)

	//Operator Aritmatika dan Boolean

	var penjumlahan = 2 - 2*3

	fmt.Println(penjumlahan)

	var firstCondition = 2 < 3
	var secondCondition = 4 > 5

	fmt.Println(firstCondition, secondCondition)

	//Operator Logika

	var wrong = false
	var right = true

	var kondisiLogikaAND = wrong && right
	fmt.Println(kondisiLogikaAND)

	var kondisiLogikaOR = wrong || right
	fmt.Println(kondisiLogikaOR)

	var kondisiLogikaNOT = !right
	fmt.Println(kondisiLogikaNOT)
}
