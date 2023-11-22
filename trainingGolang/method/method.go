package main

import "fmt"

func main() {
	mainUntukMethod()
	mainUntukMethodUsage()
	mainUntukMethodUsage2()
}

// contoh sederhana
type Person struct {
	name string
	age  int
}

func (p Person) Introduce(msg string) string {
	return fmt.Sprintf("%s My name is %s and I'm %d years old", msg, p.name, p.age)
}

func mainUntukMethod() {
	var person = Person{name: "Airell", age: 23}
	fmt.Println(person.Introduce("Hello everyone!!"))
}

// contoh lain penggunaan method
// menggunakan struct Person yang sama
func (p Person) ChangeName1() {
	p.name = "Mailo"
}
func (p *Person) ChangeName2() {
	p.name = "Mailo"
}
func mainUntukMethodUsage() {
	var person = Person{name: "Airell", age: 23}

	person.ChangeName1()
	fmt.Println("Change name with ChangeNamel method", person.name)

	person.ChangeName2()
	fmt.Println("Change name with ChangeName2 method", person.name)
}

// contoh lain penggunaan method
// menggunakan struct Person yang sama
func (p Person) Greet() {
	fmt.Println("Halo Semuanya")
}

func mainUntukMethodUsage2() {
	// Mengakses method pointer dari struct biasa
	var person1 = Person{name: "Airell", age: 23}
	person1.Greet()
	//Mengakses method pointer dari struct pointer
	var person2 = &Person{name: "Mailo", age: 15}
	person2.Greet()
}
