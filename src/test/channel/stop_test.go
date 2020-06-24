package channel

import (
	"fmt"
	"testing"
	"time"
)

func doStopSingle(ch chan int) {
	for i := 0; i < 10; i++ {
		select {
		case <-ch:
			fmt.Println("stop")
			return
		default:
			fmt.Println("put", i)
			time.Sleep(10)
		}
	}
}

func TestStopChannel(t *testing.T) {
	var ch = make(chan int)
	go doStopSingle(ch)
	time.Sleep(80)
	ch <- 1
	time.Sleep(200)
}
