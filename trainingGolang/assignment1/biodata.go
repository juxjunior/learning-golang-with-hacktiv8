// Author: Imanuel Juan Junior
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"trainingGolang/assignment1/method"
	"trainingGolang/assignment1/model"
)

func main() {
	fmt.Print("Masukkan nama: ")

	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Salah inputannya:", err)
		return
	}

	userInput = strings.TrimSpace(userInput)

	fmt.Println("Anda memasukkan:", userInput)

	peserta := method.CustomPeserta{Peserta: model.Peserta{Nama: userInput}}

	peserta.FindPeserta()
}
