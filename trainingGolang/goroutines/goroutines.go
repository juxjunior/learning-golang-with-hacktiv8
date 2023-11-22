package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go goroutine()
	mainUntukAsync()
	mainUntukAsync2()
}

// Contoh sederhana Goroutine

func goroutine() {
	fmt.Println("Hello")
}

// Contoh Goroutines (Async process #1)

func mainUntukAsync() {
	fmt.Println("main execution started")
	go firstProcess(8)

	secondProcess(8)

	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())

	fmt.Println("main execution ended")
}

func firstProcess(index int) {
	fmt.Println("First process func started")
	for i := 1; i <= index; i++ {
		fmt.Println("¡=", i)
	}
	fmt.Println("First process func ended")
}

func secondProcess(index int) {
	fmt.Println("Second process func started")
	for j := 1; j <= index; j++ {
		fmt.Println("j=", j)
	}
	fmt.Println("Second process func ended")
}

// Contoh Goroutines (Async process #2)

func mainUntukAsync2() {
	fmt.Println("main execution started")
	go firstProcess2(8)

	secondProcess2(8)

	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())

	time.Sleep(time.Second * 2)

	fmt.Println("main execution ended")
}
func firstProcess2(index int) {
	fmt.Println("First process func started")
	for i := 1; i <= index; i++ {
		fmt.Println("¡=", i)
	}
	fmt.Println("First2 process func ended")
}
func secondProcess2(index int) {
	fmt.Println("Second process func started")
	for j := 1; j <= index; j++ {
		fmt.Println("j=", j)
		fmt.Println("Second2 process func ended")
	}
}
