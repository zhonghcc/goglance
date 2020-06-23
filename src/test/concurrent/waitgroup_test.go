package concurrent

import (
	"sync"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	var sum int = 0
	lock := new(sync.Mutex)
	wait := new(sync.WaitGroup)
	wait.Add(5000)
	for i := 0; i < 5000; i++ {
		go func() {
			defer lock.Unlock()
			defer wait.Done()
			lock.Lock()
			sum++
		}()
	}
	wait.Wait()
	t.Log("sum=", sum)

}
