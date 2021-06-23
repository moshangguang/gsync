package gsync

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

func TestCompareAndSwapPointer(t *testing.T) {
	type Test struct {
		Hello string
		World int
	}
	var a *Test = nil
	b := &Test{
		Hello: "hello",
		World: 6,
	}
	ok := atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&a)), nil, unsafe.Pointer(b))
	t.Log(ok)
	t.Log(a == b)
}
func TestMutex_Lock(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(100)
	mutex := sync.Mutex{}
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			mutex.Lock()
			defer mutex.Unlock()
			time.Sleep(time.Second)
		}()
	}
	wg.Wait()
}
