package bangunruang

import (
	"bufio"
	"fmt"
	"strconv"

	"go-cli/trim"
)

func CalculateCube(reader *bufio.Reader) {
	var side float64
	var err error

	for {
		fmt.Print("Masukkan panjang sisi kubus: ")
		sideInput, _ := reader.ReadString('\n')
		sideInput = trim.TrimNewline(sideInput)

		side, err = strconv.ParseFloat(sideInput, 64)
		if err == nil {
			break
		}
		fmt.Println("Error: Masukan harus berupa angka.")
	}

	volume := side * side * side
	fmt.Printf("Volume kubus: %.2f\n", volume)
}
