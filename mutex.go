package gsync

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

type Mutex struct {
	state  *int32
	header *node
	tail   *node
}

func (m *Mutex) initQueue() {
	n := new(node)
	if atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&m.header)), nil, unsafe.Pointer(n)) {
		m.tail = n
	}
}
func (m *Mutex) addWaiter() {
	n := &node{
		ch: make(chan struct{}),
	}
	for {
		oldTail := m.tail
		if oldTail != nil {
			n.setPrevRelaxed(oldTail)
		} else {
			m.initQueue()
		}
	}
}
func (m *Mutex) Lock() {
	ok := atomic.CompareAndSwapInt32(m.state, 0, 1)
	if ok {
		return
	}
}
func (m *Mutex) Unlock() {
	s := atomic.LoadInt32(m.state)
	if s-1 < 0 {
		return
	}
	if ok := atomic.CompareAndSwapInt32(m.state, 1, 0); !ok {
		return
	}

}

func AA() {
	m := sync.Mutex{}
	m.Lock()
}
