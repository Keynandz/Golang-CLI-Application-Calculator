package bangunruang

import (
	"bufio"
	"fmt"
	"go-cli/trim"
	"strconv"
)

func CalculateRectangle(reader *bufio.Reader) {
	var length, width float64
	var err error

	// Membaca input panjang persegi panjang dan melakukan validasi
	for {
		fmt.Print("Masukkan panjang persegi panjang: ")
		lengthInput, _ := reader.ReadString('\n')
		lengthInput = trim.TrimNewline(lengthInput)

		length, err = strconv.ParseFloat(lengthInput, 64)
		if err == nil {
			break
		}
		fmt.Println("Error: Masukan harus berupa angka.")
	}

	// Membaca input lebar persegi panjang dan melakukan validasi
	for {
		fmt.Print("Masukkan lebar persegi panjang: ")
		widthInput, _ := reader.ReadString('\n')
		widthInput = trim.TrimNewline(widthInput)

		width, err = strconv.ParseFloat(widthInput, 64)
		if err == nil {
			break
		}
		fmt.Println("Error: Masukan harus berupa angka.")
	}

	// Menghitung luas persegi panjang
	area := length * width

	// Menghitung keliling persegi panjang
	perimeter := 2 * (length + width)

	fmt.Printf("Luas persegi panjang: %.2f\n", area)
	fmt.Printf("Keliling persegi panjang: %.2f\n", perimeter)
}
