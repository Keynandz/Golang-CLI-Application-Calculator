package bangunruang

import (
	"bufio"
	"fmt"
	"strconv"

	"go-cli/trim"
)

func CalculateCuboid(reader *bufio.Reader) {
	var length, width, height float64
	var err error

	for {
		fmt.Print("Masukkan panjang balok: ")
		lengthInput, _ := reader.ReadString('\n')
		lengthInput = trim.TrimNewline(lengthInput)

		length, err = strconv.ParseFloat(lengthInput, 64)
		if err == nil {
			break
		}
		fmt.Println("Error: Masukan harus berupa angka.")
	}

	for {
		fmt.Print("Masukkan lebar balok: ")
		widthInput, _ := reader.ReadString('\n')
		widthInput = trim.TrimNewline(widthInput)

		width, err = strconv.ParseFloat(widthInput, 64)
		if err == nil {
			break
		}
		fmt.Println("Error: Masukan harus berupa angka.")
	}

	for {
		fmt.Print("Masukkan tinggi balok: ")
		heightInput, _ := reader.ReadString('\n')
		heightInput = trim.TrimNewline(heightInput)

		height, err = strconv.ParseFloat(heightInput, 64)
		if err == nil {
			break
		}
		fmt.Println("Error: Masukan harus berupa angka.")
	}

	volume := length * width * height
	fmt.Printf("Volume balok: %.2f\n", volume)
}