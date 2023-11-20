package main

import "fmt"

func main() {
	slice()
}

func slice() {
	//slice normal
	var fruits = []string{"apel", "pisang", "mangga"}

	_ = fruits

	fmt.Printf("%#v\n", fruits)
	//slice make
	var fruitsMake = make([]string, 3)
	_ = fruitsMake

	fruitsMake[0] = "apple"
	fruitsMake[1] = "banana"
	fruitsMake[2] = "mango"

	fmt.Printf("%#v\n", fruitsMake)

	//slice append

	var fruitsAppend = make([]string, 3)
	_ = fruitsAppend

	fruitsAppend = append(fruitsAppend, "apple", "banana", "mango")

	fmt.Printf("%#v", fruitsAppend)

	//slice append with ellipsis

	var fruits1 = []string{"apel", "naga", "nanas"}
	var fruits2 = []string{"durian", "pisang", "stroberi"}

	fruits1 = append(fruits1, fruits2...)

	fmt.Printf("\n%#v", fruits1)

	//slice copy function

	var buah1 = []string{"apel", "mangga", "naga"}
	var buah2 = []string{"durian", "kiwi", "belimbing"}

	rujak := copy(buah1, buah2)

	fmt.Println("\nBuah 1 =>", buah1)
	fmt.Println("Buah 2 =>", buah2)
	fmt.Println("Rujak Copy =>", rujak)

	//slice slicing

	var pohon1 = []string{"apel", "pisang", "mangga", "durian", "nanas"}

	var pohon2 = pohon1[1:4]
	fmt.Printf("%#v\n", pohon2)

	var pohon3 = pohon1[0:]
	fmt.Printf("%#v\n", pohon3)

	var pohon4 = pohon1[:3]
	fmt.Printf("%#v\n", pohon4)

	var pohon5 = pohon1[:]
	fmt.Printf("%#v\n", pohon5)

	//slice gabungan slicing dan append

	var cangkok = []string{"apel", "pisang", "mangga", "durian", "nanas"}

	cangkok = append(cangkok[:3], "kiwi")

	fmt.Printf("Ini cangkok %#v\n", cangkok)

	//slice backing array

	var tamanBuah1 = []string{"apel", "pisang", "mangga", "durian", "nanas"}

	var tamanBuah2 = tamanBuah1[2:4]

	tamanBuah2[0] = "kiwi"

	fmt.Println("\nTaman Buah 1 = ", tamanBuah1)
	fmt.Println("\nTaman Buah 2 = ", tamanBuah2)

	//slice cap function

	var buahan1 = []string{"apel", "mangga", "durian", "pisang"}

	fmt.Println("buahan 1 cap:", cap(buahan1)) //4
	fmt.Println("buahan 1 len:", len(buahan1)) //4

	var buahan2 = buahan1[0:3]

	fmt.Println("buahan 2 cap:", cap(buahan2)) //4
	fmt.Println("buahan 2 len:", len(buahan2)) //3

	var buahan3 = buahan1[1:]

	fmt.Println("buahan 3 cap:", cap(buahan3)) //3
	fmt.Println("buahan 3 cap:", cap(buahan3)) //3

	//slice membuat backing array baru

	cars := []string{"Ford", "Honda", "Audi", "Toyota"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"
	fmt.Println("Cars :", cars)
	fmt.Println("newCars :", newCars)
}
