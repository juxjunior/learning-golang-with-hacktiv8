package main

import "fmt"

func main() {
	loopingAtauPerulangan()
}
func loopingAtauPerulangan() {

	//cara pertama
	for i := 0; i < 3; i++ {
		fmt.Println("Angka", i)
	}

	//cara kedua
	var perulangan = 0
	for perulangan < 3 {
		fmt.Println("Angka", perulangan)
		perulangan++
	}

	//cara ketiga
	var x = 0

	for {
		fmt.Println("Angka", x)

		x++
		if x == 5 {
			break
		}
	}

	//nestedloop

	for nestedPertama := 0; nestedPertama < 5; nestedPertama++ {
		for nestedKedua := nestedPertama; nestedKedua < 5; nestedKedua++ {
			fmt.Print(nestedKedua, " ")
		}
		fmt.Println()
	}
}
