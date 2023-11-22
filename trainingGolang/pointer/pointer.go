package main

import (
	"fmt"
	"strings"
)

func main() {
	mainUntukPointer()
	mainUntukPointerMemoryAddress()
	mainUntukPenggantianValueLewatPointer()
	mainUntukPointerSebagaiParameter()
}

//contoh sederhana pointer

func mainUntukPointer() {
	var firstNumber *int
	var secondNumber *int
	_, _ = firstNumber, secondNumber
}

//contoh pointer memory address

func mainUntukPointerMemoryAddress() {
	var firstNumber int = 4
	var secondNumber *int = &firstNumber
	fmt.Println("firstNumber (value)	:", firstNumber)
	fmt.Println("firstNumber (memori address) :", &firstNumber)

	fmt.Println("secondNumber (value) 11", *secondNumber)
	fmt.Println("secondNumber (memory address) :", secondNumber)
}

//contoh penggantian value lewat pointer

func mainUntukPenggantianValueLewatPointer() {
	var firstPerson string = "John"
	var secondPerson *string = &firstPerson
	fmt.Println("firstPerson (value) :", firstPerson)
	fmt.Println("firstPerson (memori address) :", &firstPerson)
	fmt.Println("secondNumber (value) :", *secondPerson)
	fmt.Println("secondNumber (memori address) :", secondPerson)

	fmt.Println("\n", strings.Repeat("#", 30), "\n")

	*secondPerson = "Doe"

	fmt.Println("firstPerson (value)	:", firstPerson)
	fmt.Println("firstPerson (memori address) :", &firstPerson)
	fmt.Println("secondNumber (value) :", *secondPerson)
	fmt.Println("secondNumber (memori address)	:", secondPerson)
}

//contoh pointer sebagai parameter

func mainUntukPointerSebagaiParameter() {
	var a int = 10

	fmt.Println("Before:", a)

	changeValue(&a)

	fmt.Println("After:", a)
}

func changeValue(number *int) {
	*number = 20
}
