package main

import "fmt"

func main() {
	a := 5
	ThayDoiTrenGiaTri(a)
	fmt.Println(a)
	// output: 5, ly do a chi thay doi tren function ThayDoiTrenGiaTri

	ThayDoiTrenDiaChi(&a)
	fmt.Println(a)
	// output: 200, vi thay doi persistent tren o nho roi

}

func ThayDoiTrenGiaTri(a int) {
	a = 100
}

// Truyen vao dia chi o nho
func ThayDoiTrenDiaChi(a *int) {
	*a = 200
}