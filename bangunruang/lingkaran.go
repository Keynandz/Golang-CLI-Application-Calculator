package bangunruang

import (
	"bufio"
	"fmt"
	"go-cli/trim"
	"math"
	"strconv"
)

func CalculateCircle(reader *bufio.Reader) {
	var radius float64
	var err error

	for {
		fmt.Print("Masukkan jari-jari lingkaran: ")
		radiusInput, _ := reader.ReadString('\n')
		radiusInput = trim.TrimNewline(radiusInput)

		radius, err = strconv.ParseFloat(radiusInput, 64)
		if err == nil {
			break
		}
		fmt.Println("Error: Masukan harus berupa angka.")
	}

	area := math.Pi * radius * radius
	circumference := 2 * math.Pi * radius

	fmt.Printf("Luas lingkaran: %.2f\n", area)
	fmt.Printf("Keliling lingkaran: %.2f\n", circumference)
}
