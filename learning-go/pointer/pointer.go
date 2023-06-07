package main

import "fmt"

func main() {
	var a = 100
	var b = &a

	fmt.Println(a)
	fmt.Println(b)

	*b = 2000

	fmt.Println(a)

}
