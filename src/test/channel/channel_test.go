package channel

import (
	"fmt"
	"testing"
	"time"
)

func doSingle(ch chan int) {
	var i int
	i = <-ch
	fmt.Println(i)
}

func TestSingleChannel(t *testing.T) {
	var ch = make(chan int)
	for i := 0; i < 10; i++ {
		go func() {
			doSingle(ch)
		}()
		ch <- i
		fmt.Println("put", i)
	}
	time.Sleep(2000)
}

func doBuffer(ch chan int) {
	var i int
	i = <-ch
	fmt.Println(i)
}

func TestBufferChannel(t *testing.T) {
	var ch = make(chan int, 1)
	for i := 0; i < 10; i++ {
		go func() {
			doBuffer(ch)
		}()
		ch <- i
		fmt.Println("put", i)
	}
	time.Sleep(2000)
}

func doMultiBuffer(ch chan int) {
	var i int
	i = <-ch
	fmt.Println(i)
}

func TestMutiBufferChannel(t *testing.T) {
	var ch = make(chan int, 4)
	for i := 0; i < 10; i++ {
		go func() {
			doMultiBuffer(ch)
		}()
		ch <- i
		fmt.Println("put", i)
	}
	time.Sleep(2000)
}
