package main

import "fmt"

func main() {
	//mainUntukChannels()
	mainUntukChannelAnonFunction()
}

// Contoh sederhana channels

/*func mainUntukChannels() {
	c := make(chan string)

	//Mengirim data kepada channel
	c <- value

	//Menerima data dari channel
	result := <- c
}
*/

// Contoh implementasi Channel
/*
func mainUntukChannels() {
	c := make(chan string)

	go introduce("Airell", c)
	go introduce("Nanda", c)
	go introduce("Mailo", c)

	msg1 := <-c
	fmt.Println(msg1)

	msg2 := <-c
	fmt.Println(msg2)

	msg3 := <-c
	fmt.Println(msg3)
	close(c)
}
func introduce(student string, c chan string) {
	result := fmt.Sprintf("Hai, my name is %s", student)

	c <- result
}
*/

// Contoh channel with anonymous function

func mainUntukChannelAnonFunction() {
	c := make(chan string)

	students := []string{"Airell", "Mailo", "Indah"}

	for _, v := range students {
		go func(student string) {
			fmt.Println("Student", student)
			result := fmt.Sprintf("Hai, my name is %s", student)
			c <- result
		}(v)

		for i := 1; i < 3; i++ {
			print(c)
		}
	}
}

func print(c chan string) {
	fmt.Println(<-c)
}
