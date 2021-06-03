// This is a PoC/illustrative code to show how to share a single integer that goes up in k6 on a
// single instance

package counter

import (
	"sync"
	"sync/atomic"

	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/counter", New())
}

var anonCounter int64

type namedCounter struct {
	idx *int64
}

type counter struct {
	named map[string]*namedCounter
	mu    sync.RWMutex
}

func (c *counter) Up() int64 {
	return atomic.AddInt64(&anonCounter, 1)
}

func (c *counter) UpNamed(name string) int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	counter, ok := c.named[name]

	if !ok {
		counter = &namedCounter{
			idx: new(int64),
		}
		c.named[name] = counter
	}

	return atomic.AddInt64(counter.idx, 1)
}

func New() *counter {
	return &counter{
		named: make(map[string]*namedCounter),
	}
}
