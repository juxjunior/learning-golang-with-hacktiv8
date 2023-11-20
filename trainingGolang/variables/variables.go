package main

import "fmt"

func main() {
	variables()
}

func variables() {
	// deklarasi variable
	var kalimatBebas string = "Ini isi dari String kalimatBebas ya Pak"
	var angkaBebas int = 12

	fmt.Println(kalimatBebas, angkaBebas)

	// assignment variable
	var nama string
	var nilai int

	nama = "Jux\n"
	nilai = 0

	fmt.Println(nama, nilai)

	var (
		db_password string = "xxx"
		db_name     string = "pegadaian"
		db_port     string = "8080"
		db_host     string = "localhost"
	)

	fmt.Println(db_password, db_name, db_port, db_host)

	var (
		db_password2 string
		db_name2     string
		db_port2     string
		db_host2     string
	)

	db_password2 = "xxx"
	db_name2 = "pegadaian"
	db_port2 = "8080"
	db_host2 = "localhost"

	fmt.Println(db_password2, db_name2, db_port2, db_host2)

	namaSaya := "Jux Without Data Type"

	fmt.Println(namaSaya)

	//deklarasi multi variable

	var buah, sayur, daging string
	var angkaPertama, angkaKedua, angkaKetiga = 1, 2, 3

	var namaOrang, umurOrang, alamatOrang = "Jux", 25, "Jakarta"

	buah = "Durian"
	sayur = "Sawi"
	daging = "Sapi"

	fmt.Println(buah, sayur, daging)
	fmt.Println(angkaPertama, angkaKedua, angkaKetiga)
	fmt.Println(namaOrang, umurOrang, alamatOrang)
	fmt.Printf("Cobain verbnya bosku ini type dari buah yakni %T, ini value dari namaOrang yakni %s, ini value dari angkaPertama yakni %d", buah, namaOrang, angkaPertama)
}
