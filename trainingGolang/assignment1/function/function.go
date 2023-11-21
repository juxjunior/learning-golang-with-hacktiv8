package function

import (
	"trainingGolang/assingment1/model"
)

func DataPeserta() []model.Peserta {
	pesertaList := []model.Peserta{
		{
			Nama:      "John",
			Alamat:    "Jakarta",
			Pekerjaan: "Petinju",
			Alasan:    "Null",
		},
		{
			Nama:      "Jun",
			Alamat:    "Bogor",
			Pekerjaan: "QA",
			Alasan:    "Null",
		},
		{
			Nama:      "Yoyo",
			Alamat:    "Depok",
			Pekerjaan: "Tentara",
			Alasan:    "Null",
		},
		{
			Nama:      "Japra",
			Alamat:    "Bambangan",
			Pekerjaan: "Guide",
			Alasan:    "Null",
		},
	}
	return pesertaList
}
