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
	mainUntukCalculate2()
	mainUntukPrint()
	mainUntukSum()
	mainUntukProfile()
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

func mainUntukCalculate2() {
	var diameter2 float64 = 15
	var area, circumference = calculate2(diameter2)

	fmt.Println("Area:", area)
	fmt.Println("Circumference:", circumference)
}

func calculate2(d float64) (area float64, circumference float64) {
	//Menghitung luas
	area = math.Pi * math.Pow(d/2, 2)
	//Menghitung keliling
	circumference = math.Pi * d

	return
}

// contoh menggunakan varadic function #1
func mainUntukPrint() {
	studentLists := print("John", "Ethan", "Hunt", "Wick", "Claude")
	fmt.Printf("%#v", studentLists)
}

func print(names ...string) []map[string]string {
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student%d:", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}
	return result
}

// contoh menggunakan varadic function #2
func mainUntukSum() {
	numberLists := []int{1, 2, 3, 4, 5, 6, 7, 8}

	result := sum(numberLists...)
	fmt.Println("\nResult:", result)
}

func sum(numbers ...int) int {
	total := 0

	for _, v := range numbers {
		total += v
	}
	return total
}

// contoh menggunakan varadic function #3
func mainUntukProfile() {
	profile("Udang", "Geprek", "Sambal", "Pedas", "Pete")
}

func profile(name string, favFoods ...string) {
	mergeFavFoods := strings.Join(favFoods, ",")

	fmt.Println("\nHalo, halo Saya", name)
	fmt.Println("\nSaya suka makan", mergeFavFoods)
}
