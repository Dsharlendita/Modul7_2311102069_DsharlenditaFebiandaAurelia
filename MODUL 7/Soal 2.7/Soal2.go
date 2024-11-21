package main

import (
	"fmt"
	"math"
)

func rataRata(arr []int) float64 {
	total := 0
	for _, val := range arr {
		total += val
	}
	return float64(total) / float64(len(arr))
}

func standarDeviasi(arr []int) float64 {
	mean := rataRata(arr)
	var totalDeviasi float64
	for _, val := range arr {
		totalDeviasi += math.Pow(float64(val)-mean, 2)
	}
	return math.Sqrt(totalDeviasi / float64(len(arr)))
}

func frekuensi(arr []int, target int) int {
	count := 0
	for _, val := range arr {
		if val == target {
			count++
		}
	}
	return count
}

func hapusElemen(arr []int, indeks int) []int {
	return append(arr[:indeks], arr[indeks+1:]...)
}

func main() {
	var n, x, target, indeks int

	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println("\nIsi array:")
	fmt.Println(arr)
	fmt.Println("\nElemen dengan indeks ganjil:")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
	fmt.Println("\nElemen dengan indeks genap:")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
	fmt.Print("\nMasukkan bilangan x untuk kelipatan indeks: ")
	fmt.Scan(&x)
	fmt.Println("Elemen dengan indeks kelipatan", x, ":")
	for i := 0; i < len(arr); i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
	fmt.Print("\nMasukkan indeks elemen yang ingin dihapus: ")
	fmt.Scan(&indeks)
	if indeks >= 0 && indeks < len(arr) {
		arr = hapusElemen(arr, indeks)
		fmt.Println("Array setelah elemen dihapus:")
		fmt.Println(arr)
	} else {
		fmt.Println("Indeks tidak valid!")
	}
	if len(arr) > 0 {
		fmt.Printf("\nRata-rata elemen array: %.2f\n", rataRata(arr))
		fmt.Printf("Standar deviasi elemen array: %.2f\n", standarDeviasi(arr))
	} else {
		fmt.Println("\nArray kosong, tidak dapat menghitung rata-rata dan standar deviasi.")
	}

	fmt.Print("\nMasukkan bilangan untuk menghitung frekuensi: ")
	fmt.Scan(&target)
	fmt.Printf("Frekuensi bilangan %d dalam array: %d\n", target, frekuensi(arr, target))
}
