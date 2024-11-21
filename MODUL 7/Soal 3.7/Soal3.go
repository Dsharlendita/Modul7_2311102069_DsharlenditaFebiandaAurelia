package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var daftarPemenang []string

	fmt.Print("Klub A: ")
	fmt.Scanln(&klubA)

	fmt.Print("Klub B: ")
	fmt.Scanln(&klubB)

	pertandingan := 1

	for {
		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scanln(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			fmt.Println("Pertandingan selesai lebih awal.")
			break
		}

		if skorA > skorB {
			daftarPemenang = append(daftarPemenang, klubA)
		} else if skorB > skorA {
			daftarPemenang = append(daftarPemenang, klubB)
		} else {
			daftarPemenang = append(daftarPemenang, "Draw")
		}

		pertandingan++
	}
	fmt.Println("\nHasil Pertandingan:")
	for i := 0; i < len(daftarPemenang); i++ {
		fmt.Printf("Hasil %d : %s\n", i+1, daftarPemenang[i])
	}
	jumlahKemenangan := make(map[string]int)
	for i := 0; i < len(daftarPemenang); i++ {
		hasil := daftarPemenang[i]
		if hasil != "Draw" {
			jumlahKemenangan[hasil]++
		}
	}
	fmt.Println("\nDaftar klub yang memenangkan pertandingan:")
	for klub, jumlah := range jumlahKemenangan {
		fmt.Printf("%s menang sebanyak %d kali\n", klub, jumlah)
	}

	if len(jumlahKemenangan) == 0 {
		fmt.Println("Tidak ada klub yang menang.")
	}
}
