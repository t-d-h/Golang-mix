package main

import ("fmt"
		"time"
)
// dung channel de sync giua cac goroutine, thay vi dung syncgroup
func main() {
	done := make(chan bool)
	fmt.Println("Main is going to call hellos")
	go hello(done)
	<-done
	fmt.Println("Main received data")
}

func hello(done chan bool) {
	fmt.Println("hello routine is goint to sleep for 4 secs")
	time.Sleep(4*time.Second)
	fmt.Println("waked up")
	done <- true   //ngoai ra co the dung close(done), same result
}

