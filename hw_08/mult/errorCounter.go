package mult

import "sync"

type ErrorCounter struct {
	sync.RWMutex
	value int
}

func (counter *ErrorCounter) Increase() {
	counter.Lock()
	counter.value++
	counter.Unlock()
}

func (counter *ErrorCounter) Value() int {
	counter.RLock()
	defer counter.RUnlock()
	return counter.value
}
