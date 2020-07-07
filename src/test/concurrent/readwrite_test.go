package concurrent

import (
	"strconv"
	"sync"
	"testing"
)

var rwlock *sync.RWMutex = new(sync.RWMutex)

//如果不用rwlock 会出现fatal error: concurrent map writes
func putValue(m *map[string]string, key string, value string) {
	rwlock.Lock()
	defer func() {
		rwlock.Unlock()
	}()
	(*m)[key] = value
}

func TestPutValue(t *testing.T) {
	var m map[string]string = map[string]string{}
	var group *sync.WaitGroup = new(sync.WaitGroup)
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func(i int) {
			putValue(&m, strconv.Itoa(i), strconv.Itoa(i))
			defer func() {
				group.Done()
			}()
		}(i)
	}
	group.Wait()
	t.Log(len(m))
}
