package channel

import (
	"fmt"
	"testing"
	"time"
)

func doTimeout(ch chan int) {
	var i int
	select {
	case i = <-ch:
		fmt.Println(i)

	case <-time.After(100):
		fmt.Println("timeout")
	}
}

func TestTimeoutChannel(t *testing.T) {
	var ch = make(chan int, 1)
	for i := 0; i < 10; i++ {
		go doTimeout(ch)
	}
	for i := 0; i < 10; i++ {
		if i == 8 {
			time.Sleep(2500)
		} else {
			ch <- i
		}
		fmt.Println("put", i)
	}
	time.Sleep(2000)
}
