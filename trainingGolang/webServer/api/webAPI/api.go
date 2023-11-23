package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strconv"
)

func main() {
	mainUntukAPIGET()
	//mainUntukAPIPOST()
}

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var employees = []Employee{
	{ID: 1, Name: "Jontor", Age: 25, Division: "Marketing"},
	{ID: 2, Name: "Joni", Age: 50, Division: "Treasury"},
	{ID: 3, Name: "Jack", Age: 30, Division: "IT"},
}

var PORT = ":8080"

func mainUntukAPIGET() {
	http.HandleFunc("/employees", getEmployees)
	mainUntukAPIPOST()
	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

// Contoh menggunakan HTML
func getEmployees(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		fp := path.Join("/PT Pegadaian/Divisi Pengembangan Aplikasi TI/Training Golang 202311/go/src/trainingGolang/webServer/api/webAPI", "template.html") //sesuaikan dengan path masing-masing
		tpl, err := template.ParseFiles(fp)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tpl.Execute(w, employees)
		return
	}
	http.Error(w, "Invalid method", http.StatusBadRequest)
}

/*
Cobalah menembak request GET ini menggunakan Postman
dengan meng-import curl di bawah ini

curl --location 'http://localhost:8080/employees' /
--data ''
*/

func mainUntukAPIPOST() {
	http.HandleFunc("/employee", createEmployee)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		name := r.FormValue("name")
		age := r.FormValue("age")
		divison := r.FormValue("division")

		convertAge, err := strconv.Atoi(age)
		if err != nil {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}
		newEmployee := Employee{
			ID:       len(employees) + 1,
			Name:     name,
			Age:      convertAge,
			Division: divison,
		}

		employees = append(employees, newEmployee)

		json.NewEncoder(w).Encode(newEmployee)
		return
	}
	http.Error(w, "Invalid method", http.StatusBadRequest)
}

/*
Cobalah menembak request POST ini menggunakan Postman
dengan meng-import curl di bawah ini
curl --location 'http://localhost:8080/employee' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'name=Jux' \
--data-urlencode 'age=25' \
--data-urlencode 'division=Business'
*/
