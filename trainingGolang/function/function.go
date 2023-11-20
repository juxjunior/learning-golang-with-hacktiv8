package main

import "fmt"

func main() {
	greet("John", 23)
	salam("Jux", "Jakpus")
}

func greet(name string, age int8) {
	fmt.Printf("Hello there! My name is %s and I'm %d years old", name, age)
}

func salam(name, address string) {
	fmt.Println("Hello there! My name is", name)
	fmt.Println("I live in ", address)
}
