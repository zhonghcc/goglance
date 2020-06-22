package concurrent

import (
	"testing"
	"time"
)

func TestGo(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			t.Log(i)
		}(i)
	}
	time.Sleep(100)
}

func TestGoWrong(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			t.Log(i)
		}()
	}
	time.Sleep(1000)
}
