package main

import (
	"fmt"
	"strings"
)

func main() {
	array()
}

func array() {
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}
	fmt.Println(numbers)

	var kalimat = [3]string{"Aku", "Dia", "Kamu"}
	fmt.Println(kalimat)

	var buah = [3]string{"durian", "naga", "mangga"}
	buah[0] = "apel"
	buah[1] = "pisang"
	buah[2] = "kiwi"

	fmt.Printf("%#v\n", buah)

	var fruits = [3]string{"apel", "pisang", "mangga"}
	for i, v := range fruits {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}

	fmt.Println(strings.Repeat("#", 25))

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Index: %d, Value: %s\n", i, fruits[i])
	}

	fmt.Println(strings.Repeat("#", 25))

	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}
