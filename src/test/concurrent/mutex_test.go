package concurrent

import (
	"sync"
	"testing"
	"time"
)

func TestAddWrong(t *testing.T) {
	var sum int = 0
	for i := 0; i < 5000; i++ {
		go func() {
			sum++
		}()
	}
	time.Sleep(3000)
	t.Log("sum=", sum)
}

func TestMutex(t *testing.T) {
	var sum int = 0
	lock := new(sync.Mutex)
	for i := 0; i < 5000; i++ {
		go func() {
			defer lock.Unlock()
			lock.Lock()
			sum++
		}()
	}
	time.Sleep(3000)
	t.Log("sum=", sum)

}
