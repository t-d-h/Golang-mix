package main
import "fmt"


type Rectangle struct {
	height int
	wide int
}

func (r Rectangle) area() int {
	return r.height * r.wide
}

func (r Rectangle) perimeter() int {
    return 2 * (r.height + r.wide)
}

func main() {
	r := Rectangle{2,5}
	fmt.Println(r.area())
    fmt.Println(r.perimeter())
}