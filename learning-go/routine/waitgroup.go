package main
import (
	 "fmt"
	"sync"
)

func main() {
	array := []int{}
	chunk1 := 0
	chunk2 := 0
	for i := 0; i < 20; i ++ {
		array = append (array, i)
	}
	var wg sync.WaitGroup

	wg.Add(2) // cho 2 goroutine Done()
	go func() {
		defer wg.Done() // defer: chay dong nay khi function chay xong
		for i := 0; i < 10; i++ {
			chunk1 += array[i]
		}
		fmt.Println("chunk1: ", chunk1)
	}()

	go func() {
		defer wg.Done()
		for i := 10; i < 20; i++ {
			chunk2 += array[i]
		}
		fmt.Println("chunk2: ", chunk2)
	}()
	wg.Wait() //Wait "2" goroutine chay xong thi moi chay nhung dong tiep theo
	fmt.Println("Tong: ", chunk1+chunk2)

} 