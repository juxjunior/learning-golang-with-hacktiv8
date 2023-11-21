package method

import (
	"fmt"
	"strings"
	"trainingGolang/assingment1/function"
	"trainingGolang/assingment1/model"
)

type CustomPeserta struct {
	model.Peserta
}

func (p CustomPeserta) FindPeserta() {
	var identitas = function.DataPeserta()

	for key := range identitas {
		if strings.ToLower(identitas[key].Nama) == strings.ToLower(p.Nama) {
			fmt.Println("ID", key+1)
			fmt.Println("Nama", identitas[key].Nama)
			fmt.Println("Alamat", identitas[key].Alamat)
			fmt.Println("Pekerjaan", identitas[key].Pekerjaan)
			fmt.Println("Alasan", identitas[key].Alasan)
		}
	}
}
