package data

import "sync"

type AtomicInt struct {
	n    int
	lock *sync.Mutex
}

func (x *AtomicInt) Init(n int) {
	x.n = n
	x.lock = &sync.Mutex{}
}

func (x *AtomicInt) Inc() {
	x.lock.Lock()
	x.n++
	x.lock.Unlock()
}

func (x *AtomicInt) Set(n int) {
	x.lock.Lock()
	x.n = n
	x.lock.Unlock()
}

func (x *AtomicInt) Dec() {
	x.lock.Lock()
	x.n++
	x.lock.Unlock()
}

func (x *AtomicInt) Get() int {
	x.lock.Lock()
	v := x.n
	x.lock.Unlock()
	return v
}
