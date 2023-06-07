package main
import "fmt"

func main() {
	a := 1
	fmt.Printf("a truoc khi thay doi la %v\n", a)

	b := &a
	//Cach khac de khai bao con tro:
	// c:=new(int)
	// *c = 23
	change(b)
	
	fmt.Printf("a sau khi thay doi la %v\n", a)

	*b+=2
	fmt.Printf("a + 2 := %v\n", a)

}

func change(a *int) {
	*a = 2
}