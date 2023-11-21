package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	greet("John", 23)
	salam("Jux", "Jakpus")
	mainUntukSalam2()
	mainUntukCalculate()
}

// contoh sederhana
func greet(name string, age int8) {
	fmt.Printf("Hello there! My name is %s and I'm %d years old", name, age)
}

// contoh menggunakan paramater
func salam(name, address string) {
	fmt.Println("Hello there! My name is", name)
	fmt.Println("I live in ", address)
}

// contoh menggunakan return
func mainUntukSalam2() {
	var names = []string{"John", "Jux"}
	var printMsg = salam2("Halo bro", names)
	fmt.Println(printMsg)
}
func salam2(msg string, names []string) string {
	var joinStr = strings.Join(names, " ")

	var result string = fmt.Sprintf("%s %s", msg, joinStr)

	return result
}

// contoh menggunakan return multiple values
func mainUntukCalculate() {
	var diameter float64 = 15
	var luas, keliling = calculate(diameter)

	fmt.Println("Luas:", luas)
	fmt.Println("Keliling:", keliling)
}

func calculate(d float64) (float64, float64) {
	//Menghitung luas
	var luas float64 = math.Pi * math.Pow(d/2, 2)
	//Menghitung keliling
	var keliling = math.Pi * d

	return luas, keliling
}

// contoh menggunakan function predefined return values
