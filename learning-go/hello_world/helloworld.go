package main
import "fmt"

func main() {
	fmt.Println("Hello Khoai!")

	var a = []interface{}{123, "abc"}

	fmt.Println(a...)
	fmt.Println(a)
	fmt.Println(b)
}