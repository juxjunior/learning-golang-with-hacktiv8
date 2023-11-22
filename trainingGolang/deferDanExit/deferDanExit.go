package main

import (
	"fmt"
	"os"
)

func main() {
	mainUntukDefer()
	mainUntukDefer2()
	mainUntukExit()
}

// Contoh defer sederhana

func mainUntukDefer() {
	defer fmt.Println("defer function starts to execute")
	fmt.Println("Hai everyone")
	fmt.Println("Welcome back to Go learning center")
}

// Contoh lain penggunaan defer

func mainUntukDefer2() {
	callDeferFunc()
	fmt.Println("Hai everyone 2")
}

func callDeferFunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("defer function starts to execute 2")
}

// Contoh exit

func mainUntukExit() {
	defer fmt.Println("Invoke with defer")
	fmt.Println("Before Exiting")
	os.Exit(1)
}
