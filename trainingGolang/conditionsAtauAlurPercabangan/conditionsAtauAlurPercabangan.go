package main

import "fmt"

func main() {
	conditionsAtauAlurPercabangan()
}

func conditionsAtauAlurPercabangan() {
	//ifelse

	var currentYear = 2023

	if age := currentYear - 1997; age <= 10 {
		fmt.Println("Umur di bawah 10 Tahun")
	} else {
		fmt.Println("Umur di atas 10 Tahun")
	}

	//Switchcase
	var score = 8

	switch score {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("Good")
	default:
		fmt.Println("Not Bad")
	}

	//Switch With Relational Operators
	var poin = 7

	switch {
	case poin == 8:
		fmt.Println("Perfect")
	case poin > 5:
		fmt.Println("Good")
	default:
		fmt.Println("Not Bad")
	}

	//Switch (fallthrough keyword)

}
