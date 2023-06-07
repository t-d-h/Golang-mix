package main
import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Application start")

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutines: ", i)
		}
	}()

	fmt.Println("Application end")
	// cho` function tren ket thuc, neu khong sleep thi no se "khai hoa cai function thoi xong no chay tiep code o duoi"
	// neu func main chay xong ma goroutine chua chay xong thi no se huy cai goroutine do
	time.Sleep(100 * time.Second)
}