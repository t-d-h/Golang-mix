package main
import (
	"fmt"
	"time"	
)
// dung go de gui du lieu giua 2 goroutines
func main() {
	channelll := make(chan bool)
	fmt.Println("Start")
	go count(channelll)
	
	<-channelll 
	
}

func count(channelll chan bool) {
	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		fmt.Println("Goroutines: ", i)
        
    }
	channelll <- true
}