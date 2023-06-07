package main
import "fmt"

type empl struct {
	name string
	age int
}

func (e empl) changeName(val string) { //value receiver
	e.name = val
}

func (e *empl) changeAge(val int) { // poitner receiver
	e.age = val
}

func main() {
	empl1 := empl {"hoan", 22}
	empl1.changeName("Tam")
	fmt.Printf("Name after change: %v\n", empl1.name)

	empl1.changeAge(24)
	fmt.Printf("Age after change: %v\n", empl1.age)
}
/* conclusion:
- Dung poiter receiver de thay doi gia tri receiver ben trong method (thay doi thuoc tinh cua emlp1)

- Khi chi muon truy cap gia tri cua empl1 thi dung value receiver 