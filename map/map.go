package Map

import (
	"sync"
)

type Map struct {
	sync.RWMutex
	M map[interface{}]interface{}
}

func NewMap() *Map {
	return &Map{M: make(map[interface{}]interface{})}
}

func (this *Map) Get(key interface{}) (interface{}, bool) {
	this.RLock()
	val, ok := this.M[key]
	this.RUnlock()
	return val, ok
}

func (this *Map) Set(key interface{}, val interface{}) {
	this.Lock()
	this.M[key] = val
	this.Unlock()
}

func (this *Map) Len() int {
	this.RLock()
	length := len(this.M)
	this.RUnlock()
	return length
}

func (this *Map) Delete(key interface{}) {
	this.Lock()
	delete(this.M, key)
	this.Unlock()
}

func (this *Map) BatchDelete(keys []interface{}) {
	count := len(keys)
	if count == 0 {
		return
	}

	this.Lock()
	for i := 0; i < count; i++ {
		delete(this.M, keys[i])
	}
	this.Unlock()
}
