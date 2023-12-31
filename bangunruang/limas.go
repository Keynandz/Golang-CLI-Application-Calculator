package bangunruang

import (
	"bufio"
	"fmt"
	"strconv"

	"go-cli/trim"
)

func CalculatePyramid(reader *bufio.Reader) {
	var baseArea, height float64
	var err error

	// Membaca input luas alas limas dan melakukan validasi
	for {
		fmt.Print("Masukkan luas alas limas: ")
		baseAreaInput, _ := reader.ReadString('\n')
		baseAreaInput = trim.TrimNewline(baseAreaInput)

		baseArea, err = strconv.ParseFloat(baseAreaInput, 64)
		if err == nil {
			break
		}
		fmt.Println("Error: Masukan harus berupa angka.")
	}

	// Membaca input tinggi limas dan melakukan validasi
	for {
		fmt.Print("Masukkan tinggi limas: ")
		heightInput, _ := reader.ReadString('\n')
		heightInput = trim.TrimNewline(heightInput)

		height, err = strconv.ParseFloat(heightInput, 64)
		if err == nil {
			break
		}
		fmt.Println("Error: Masukan harus berupa angka.")
	}

	// Menghitung volume limas
	volume := (1.0 / 3.0) * baseArea * height
	fmt.Printf("Volume limas: %.2f\n", volume)
}
