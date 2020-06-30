package concurrent

import (
	"fmt"
	"sync"
	"testing"
)

var once sync.Once

func initOnce() {
	fmt.Println("init once")
}

func TestOnce(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(initOnce)
			fmt.Println("do ", i)
		}()
	}
}
