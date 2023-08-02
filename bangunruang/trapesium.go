package bangunruang

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"go-cli/trim"
)

func CalculateTrapezoid(reader *bufio.Reader) {
	var top, bottom, height float64
	var err error

	for {
		fmt.Print("Masukkan panjang sisi atas trapesium: ")
		topInput, _ := reader.ReadString('\n')
		topInput = trim.TrimNewline(topInput)

		top, err = strconv.ParseFloat(topInput, 64)
		if err == nil {
			break
		}
		fmt.Println("Error: Masukan harus berupa angka.")
	}

	for {
		fmt.Print("Masukkan panjang sisi bawah trapesium: ")
		bottomInput, _ := reader.ReadString('\n')
		bottomInput = trim.TrimNewline(bottomInput)

		bottom, err = strconv.ParseFloat(bottomInput, 64)
		if err == nil {
			break
		}
		fmt.Println("Error: Masukan harus berupa angka.")
	}

	for {
		fmt.Print("Masukkan tinggi trapesium: ")
		heightInput, _ := reader.ReadString('\n')
		heightInput = trim.TrimNewline(heightInput)

		height, err = strconv.ParseFloat(heightInput, 64)
		if err == nil {
			break
		}
		fmt.Println("Error: Masukan harus berupa angka.")
	}

	area := 0.5 * (top + bottom) * height
	perimeter := top + bottom + 2*math.Sqrt((top-bottom)*(top-bottom)/4+height*height)

	fmt.Printf("Luas trapesium: %.2f\n", area)
	fmt.Printf("Keliling trapesium: %.2f\n", perimeter)
}