package concurrent

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestAtomic(t *testing.T) {
	var sum int32 = 0
	for i := 0; i < 5000; i++ {
		go func() {
			atomic.AddInt32(&sum, 1)
		}()
	}
	time.Sleep(3000)
	t.Log("sum=", sum)

}
