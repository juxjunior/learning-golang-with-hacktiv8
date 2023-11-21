package main

import (
	"fmt"
	"strings"
)

func main() {
	mainUntukValueKeStruct()
	mainUntukInitializingStruct()
	mainUntukPointerToStruct()
	mainUntukEmbededStruct()
	mainUntukAnonymousStruct()
}

// contoh struct sederhana

type Employee struct {
	name     string
	age      int
	division string
}

// contoh struct (memberikan value ke struct)

func mainUntukValueKeStruct() { //menggunakan struct Employee di atas
	var employee Employee
	employee.name = "Airell"
	employee.age = 23
	employee.division = "Curriculum Developer"
	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.division)
}

// contoh menginisialisasi struct

func mainUntukInitializingStruct() {
	var employee1 = Employee{}
	employee1.name = "Airell"
	employee1.age = 23
	employee1.division = "Curriculum Developer"
	var employee2 = Employee{name: "Ananda", age: 23, division: "Finance"}
	fmt.Printf("Employeel: %+v\n", employee1)
	fmt.Printf("Employee2: %+v\n", employee2)
}

// contoh pointer to struct

func mainUntukPointerToStruct() {
	var employee1 = Employee{name: "Airell", age: 23, division: "Curriculum Developer"}
	var employee2 *Employee = &employee1
	fmt.Println("Employeel name:", employee1.name)
	fmt.Println("Employee2 name", employee2.name)

	fmt.Println(strings.Repeat("#", 20))
	employee2.name = "Ananda"
	fmt.Println("Employeel name:", employee1.name)
	fmt.Println("Employee2 name", employee2.name)
}

// contoh embeded struct

type Person struct {
	name string
	age  int
}

type Karyawan struct {
	division string
	person   Person
}

func mainUntukEmbededStruct() {
	var employee1 = Karyawan{}
	employee1.person.name = "Airell"
	employee1.person.age = 23
	employee1.division = "Curriculum Developer"

	fmt.Printf("%+v", employee1)
}

// contoh anonymous struct

func mainUntukAnonymousStruct() {
	//Anonymous struct tapa pengisian field
	var employee1 = struct {
		person   Person
		division string
	}{}

	employee1.person = Person{name: "Airell", age: 23}
	employee1.division = "Curriculum developer"
	//Anonymous struct dengan pengisian field
	var employee2 = struct {
		person   Person
		division string
	}{
		person:   Person{name: "Ananda", age: 23},
		division: "Finance",
	}
	fmt.Printf("\nEmployee1: %+v\n", employee1)
	fmt.Printf("Employee1: %+v\n", employee2)
}
