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
		// Menampilkan menu pilihan bangun ruang
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
			// Memanggil fungsi untuk menghitung luas dan keliling lingkaran
			bangunruang.CalculateCircle(reader)
		case "2":
			// Memanggil fungsi untuk menghitung luas dan keliling segitiga
			bangunruang.CalculateTriangle(reader)
		case "3":
			// Memanggil fungsi untuk menghitung luas dan keliling persegi panjang
			bangunruang.CalculateRectangle(reader)
		case "4":
			// Memanggil fungsi untuk menghitung luas dan keliling persegi
			bangunruang.CalculateSquare(reader)
		case "5":
			// Memanggil fungsi untuk menghitung luas dan keliling trapesium
			bangunruang.CalculateTrapezoid(reader)
		case "6":
			// Memanggil fungsi untuk menghitung luas dan volume kubus
			bangunruang.CalculateCube(reader)
		case "7":
			// Memanggil fungsi untuk menghitung luas dan volume balok
			bangunruang.CalculateCuboid(reader)
		case "8":
			// Memanggil fungsi untuk menghitung luas dan volume limas
			bangunruang.CalculatePyramid(reader)
		case "9":
			// Menampilkan pesan dan keluar dari aplikasi
			fmt.Println("Terima kasih telah menggunakan aplikasi ini. Sampai jumpa lagi!")
			return
		default:
			// Menampilkan pesan kesalahan untuk pilihan yang tidak valid
			fmt.Println("Pilihan tidak valid. Silakan pilih 1, 2, 3, 4, 5, 6, 7, 8, dan 9.")
		}
	}
}
