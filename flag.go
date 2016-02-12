package gopter

import "sync"

type Flag struct {
	lock sync.RWMutex
	flag bool
}

func (f *Flag) Get() bool {
	f.lock.RLock()
	defer f.lock.RUnlock()
	return f.flag
}

func (f *Flag) Set() {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.flag = true
}

func (f *Flag) Unset() {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.flag = false
}
