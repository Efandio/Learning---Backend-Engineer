package main

import (
	"fmt"
)




func main() {

	fmt.Println("Penggunaan Go Map [Hash Map]")

	var nilaiSiswa = make(map[string]int)
	fmt.Printf("Map awal: %v\n", nilaiSiswa)

	fmt.Println("Menambah nilai siswa:")
	nilaiSiswa["Doni"] = 95
	nilaiSiswa["Andi"] = 80
	nilaiSiswa["Sila"] = 65
	nilaiSiswa["Indi"] = 90
	nilaiSiswa["Deni"] = 100

	fmt.Printf("Setelah ditambah %v\n", nilaiSiswa)

	fmt.Println("Mengambil data...")

	nilaiDoni := nilaiSiswa["Doni"]
	fmt.Printf("Nilainya adalah %v\n", nilaiDoni)

	nilaiDavid := nilaiSiswa["David"]
	fmt.Printf("Nilainya adalah %v\n",  nilaiDavid)

	fmt.Println("Mengambil data dengan cek keberadaan")
	score, exist := nilaiSiswa["Andi"]
	if exist {
		fmt.Printf("Nilai Andi adalah %v\n", score)
	} else {
		fmt.Printf("Nilai Andi %v\n", score)
	}

	fmt.Println("Menghapus data")
	delete(nilaiSiswa, "Doni")
	fmt.Printf("Nilai siswa saat ini %v\n", nilaiSiswa)
	delete(nilaiSiswa, "Botol")
	fmt.Printf("Nilai siswa saat ini %v\n", nilaiSiswa)

	for nama, nilai := range nilaiSiswa {
		fmt.Printf("Nama siswa: %v, nilainya: %v ", nama, nilai)
	}
}
