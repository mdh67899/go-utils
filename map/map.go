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
	defer this.RUnlock()
	val, ok := this.M[key]
	return val, ok
}

func (this *Map) Set(key interface{}, val interface{}) {
	this.Lock()
	defer this.Unlock()
	this.M[key] = val
}

func (this *Map) Len() int {
	this.RLock()
	defer this.RUnlock()
	return len(this.M)
}

func (this *Map) Delete(key interface{}) {
	this.Lock()
	defer this.Unlock()
	delete(this.M, key)
}

func (this *Map) BatchDelete(keys []interface{}) {
	count := len(keys)
	if count == 0 {
		return
	}

	this.Lock()
	defer this.Unlock()
	for i := 0; i < count; i++ {
		delete(this.M, keys[i])
	}
}
