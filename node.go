package gsync

import (
	"sync/atomic"
	"unsafe"
)

type node struct {
	ch   chan struct{}
	prev *node
	next *node
}

func (node *node) setPrevRelaxed(prev *node) {
	atomic.SwapPointer((*unsafe.Pointer)(unsafe.Pointer(&node.prev)), unsafe.Pointer(prev))
}
