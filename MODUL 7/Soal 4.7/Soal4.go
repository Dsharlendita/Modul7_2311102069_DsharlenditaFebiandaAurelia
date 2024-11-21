package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

// Isi array dengan input dari pengguna
func isiArray(t *tabel, n *int) {
	var input string
	fmt.Print("Masukkan kata: ")
	// Membaca input sampai dengan Enter (tanpa titik)
	fmt.Scanln(&input)
	// Memasukkan karakter satu per satu ke dalam array
	for i := 0; i < len(input); i++ {
		(*t)[*n] = rune(input[i]) // Konversi karakter menjadi rune
		*n++                      // Increment n untuk menyimpan posisi berikutnya
	}
}

// Cetak isi array
func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(string(t[i])) // Menampilkan karakter sebagai string
	}
	fmt.Println()
}

// Balikkan urutan array
func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

// Memeriksa apakah array membentuk palindrom
func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	// Isi array dengan input pengguna
	isiArray(&tab, &m)

	// Tampilkan teks asli
	fmt.Print("Teks        : ")
	cetakArray(tab, m)

	// Balikkan urutan teks dan tampilkan
	balikanArray(&tab, m)
	fmt.Print("Reverse Teks: ")
	cetakArray(tab, m)

	// Cek apakah teks membentuk palindrom
	if palindrom(tab, m) {
		fmt.Println("Palindrom   ? true")
	} else {
		fmt.Println("Palindrom   ? false")
	}
}
