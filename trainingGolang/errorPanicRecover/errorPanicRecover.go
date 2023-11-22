package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	mainUntukError()
	mainUntukCustomError()
	mainUntukPanic()
	mainUntukRecover()
}

// Contoh Error sederhana

func mainUntukError() {
	var number int
	var err error

	number, err = strconv.Atoi("123GH")

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	number, err = strconv.Atoi("123")

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}
}

// Contoh Custom Error

func mainUntukCustomError() {
	var password string
	fmt.Scanln(&password)

	if valid, err := validPassword(password); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(valid)
	}
}

func validPassword(password string) (string, error) {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("password has to have more than 4 characters")
	}
	return "Valid password", nil
}

// Contoh Panic

func mainUntukPanic() {
	var password string
	fmt.Scanln(&password)
	if valid, err := validPassword(password); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}
}

// Contoh Recover

func mainUntukRecover() {
	defer catchErr()

	var password string

	fmt.Scanln(&password)

	if valid, err := validPassword(password); err != nil {

		panic(err.Error())
	} else {
		fmt.Println(valid)
	}
}

func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}
