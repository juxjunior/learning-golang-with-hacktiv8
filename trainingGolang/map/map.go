package main

import "fmt"

func main() {
	mapContoh()
}

func mapContoh() {
	//contoh sederhana

	var person = map[string]string{
		"name":    "Jux",
		"age":     "25",
		"address": "Jakarta",
	}

	fmt.Println("name:", person["name"])
	fmt.Println("age:", person["age"])
	fmt.Println("address:", person["address"])

	//looping menggunakan map

	var orang = map[string]string{
		"name":    "Jux",
		"age":     "25",
		"address": "Jakarta",
	}

	for key, value := range orang {
		fmt.Println(key, ":", value)
	}

	//map deleting value

	var manusia = map[string]string{
		"name":    "Jux",
		"age":     "25",
		"address": "Jakarta",
	}

	fmt.Println("Sebelum menghapus:", manusia)

	delete(manusia, "address")

	fmt.Println("Setelah menghapus:", manusia)

	//map detecting value

	var identitas = map[string]string{
		"name":    "Jux",
		"age":     "25",
		"address": "Jakarta",
	}

	value, exist := identitas["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value doesn't exist")
	}

	delete(identitas, "age")

	value, exist = identitas["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value has been deleted")
	}

	// Kombinasi slice dan map

	var people = []map[string]string{
		{"name": "Jux", "age", "25"},
		{"name": "John", "age", "15"},
		{"name": "Juli", "age", "21"},
	}

	for i, person := range people {
		fmt.Printf()
	}
}
