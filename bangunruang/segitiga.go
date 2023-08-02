package bangunruang

import (
	"bufio"
	"fmt"
	"go-cli/trim"
	"math"
	"strconv"
)

func CalculateTriangle(reader *bufio.Reader) {
	var base, height float64
	var err error

	// Membaca input panjang alas segitiga dan melakukan validasi
	for {
		fmt.Print("Masukkan panjang alas segitiga: ")
		baseInput, _ := reader.ReadString('\n')
		baseInput = trim.TrimNewline(baseInput)

		base, err = strconv.ParseFloat(baseInput, 64)
		if err == nil {
			break
		}
		fmt.Println("Error: Masukan harus berupa angka.")
	}

	// Membaca input tinggi segitiga dan melakukan validasi
	for {
		fmt.Print("Masukkan tinggi segitiga: ")
		heightInput, _ := reader.ReadString('\n')
		heightInput = trim.TrimNewline(heightInput)

		height, err = strconv.ParseFloat(heightInput, 64)
		if err == nil {
			break
		}
		fmt.Println("Error: Masukan harus berupa angka.")
	}

	// Menghitung luas segitiga
	area := 0.5 * base * height

	// Menghitung keliling segitiga
	perimeter := base + 2*math.Sqrt(base*base/4+height*height)

	fmt.Printf("Luas segitiga: %.2f\n", area)
	fmt.Printf("Keliling segitiga: %.2f\n", perimeter)
}
