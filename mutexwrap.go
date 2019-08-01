package log

import "sync"

type MutexWrap struct {
	mux sync.RWMutex
	off bool
}

func (mw *MutexWrap) Lock() {
	if !mw.off {
		mw.mux.Lock()
	}
}

func (mw *MutexWrap) Unlock() {
	if !mw.off {
		mw.mux.Unlock()
	}
}

func (mw *MutexWrap) RLock() {
	if !mw.off {
		mw.mux.RLock()
	}
}

func (mw *MutexWrap) RUnlock() {
	if !mw.off {
		mw.mux.RUnlock()
	}
}

func (mw *MutexWrap) NoLock(noLock bool) {
	mw.off = noLock
}
