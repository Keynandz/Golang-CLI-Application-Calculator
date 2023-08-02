package main

import (
	"bufio"
	"fmt"
	"os"

	"go-cli/bangunruang"
	"go-cli/trim"
)

func main() {
	fmt.Println("Selamat datang di aplikasi CLI sederhana untuk menghitung luas dan keliling beberapa bangun ruang!")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nPilih bangun ruang:")
		fmt.Println("1. Lingkaran")
		fmt.Println("2. Segitiga")
		fmt.Println("3. Persegi Panjang")
		fmt.Println("4. Persegi")
		fmt.Println("5. Trapesium")
		fmt.Println("6. Kubus")
		fmt.Println("7. Balok")
		fmt.Println("8. Limas")
		fmt.Println("9. Keluar")

		fmt.Print("Masukkan pilihan (1/2/3/4/5/6/7/8/9): ")
		choiceInput, _ := reader.ReadString('\n')
		choiceInput = trim.TrimNewline(choiceInput)

		switch choiceInput {
		case "1":
			bangunruang.CalculateCircle(reader)
		case "2":
			bangunruang.CalculateTriangle(reader)
		case "3":
			bangunruang.CalculateRectangle(reader)
		case "4":
			bangunruang.CalculateSquare(reader)
		case "5":
			bangunruang.CalculateTrapezoid(reader)
		case "6":
			bangunruang.CalculateCube(reader)
		case "7":
			bangunruang.CalculateCuboid(reader)
		case "8":
			bangunruang.CalculatePyramid(reader)
		case "9":
			fmt.Println("Terima kasih telah menggunakan aplikasi ini. Sampai jumpa lagi!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih 1, 2, 3, 4, 5, 6, 7, 8, dan 9.")
		}
	}
}
