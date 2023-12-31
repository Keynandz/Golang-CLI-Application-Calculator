package bangunruang

import (
	"bufio"
	"fmt"
	"go-cli/trim"
	"strconv"
)

func CalculateSquare(reader *bufio.Reader) {
	var side float64
	var err error

	// Membaca input panjang sisi persegi dan melakukan validasi
	for {
		fmt.Print("Masukkan panjang sisi persegi: ")
		sideInput, _ := reader.ReadString('\n')
		sideInput = trim.TrimNewline(sideInput)

		side, err = strconv.ParseFloat(sideInput, 64)
		if err == nil {
			break
		}
		fmt.Println("Error: Masukan harus berupa angka.")
	}

	// Menghitung luas persegi
	area := side * side

	// Menghitung keliling persegi
	perimeter := 4 * side

	fmt.Printf("Luas persegi: %.2f\n", area)
	fmt.Printf("Keliling persegi: %.2f\n", perimeter)
}
