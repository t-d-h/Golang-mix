//method la function danh cho kieu du lieu dac biet (struct or non struct), kieu du lieu nay goi la receiver

package main
import "fmt"

type employee struct {
	name string
	age int
	sal int
}

func (e employee) GetSal() {
	fmt.Printf("sal cua %v la %v", e.name, e.sal)
}

func main() {
	emp1 := employee{
		name: "hoan",
		age: 22,
		sal: 2000,
	}
	emp1.GetSal()
}