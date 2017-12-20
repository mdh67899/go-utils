package List

import (
	"container/list"
	"sync"
)

type Linklist struct {
	sync.RWMutex
	L *list.List
}

func NewLinklist() *Linklist {
	return &Linklist{L: list.New()}
}

func (this *Linklist) Init() {
	this.Lock()
	this.L = this.L.Init()
	this.Unlock()
}

func (this *Linklist) len() int {
	return this.L.Len()
}

func (this *Linklist) Len() int {
	this.RLock()
	length := this.len()
	this.RUnlock()
	return length
}

func (this *Linklist) Remove(v *list.Element) interface{} {
	this.Lock()
	value := this.L.Remove(v)
	this.Unlock()
	return value
}

func (this *Linklist) frontByNum(count int) []interface{} {
	items := make([]interface{}, 0, count)
	CurrentItem := this.L.Front()

	for i := 0; i < count; i++ {
		if CurrentItem == nil {
			return items
		}

		items = append(items, CurrentItem.Value)
		CurrentItem = CurrentItem.Next()
	}

	return items
}

func (this *Linklist) backByNum(count int) []interface{} {
	items := make([]interface{}, 0, count)
	CurrentItem := this.L.Back()

	for i := 0; i < count; i++ {
		if CurrentItem == nil {
			return items
		}

		items = append(items, CurrentItem.Value)
		CurrentItem = CurrentItem.Prev()
	}

	return items
}

func (this *Linklist) BatchFront(count int) []interface{} {
	if count <= 0 {
		return []interface{}{}
	}

	this.RLock()

	total := this.len()

	if total == 0 {
		this.RUnlock()
		return []interface{}{}
	}

	if count > total {
		count = total
	}

	items := this.frontByNum(count)

	this.RUnlock()
	return items
}

func (this *Linklist) BatchBack(count int) []interface{} {
	if count <= 0 {
		return []interface{}{}
	}

	this.RLock()

	total := this.len()

	if total == 0 {
		this.RUnlock()
		return []interface{}{}
	}

	if count > total {
		count = total
	}

	items := this.backByNum(count)

	this.RUnlock()
	return items
}

func (this *Linklist) pushFront(v interface{}) *list.Element {
	return this.L.PushFront(v)
}

func (this *Linklist) PushFront(v interface{}) *list.Element {
	this.Lock()
	e := this.pushFront(v)
	this.Unlock()
	return e
}

func (this *Linklist) BatchPushFront(v []interface{}) []*list.Element {
	length := len(v)

	if length == 0 {
		return []*list.Element{}
	}

	items := make([]*list.Element, 0, length)

	this.Lock()
	for index, _ := range v {
		items = append(items, this.pushFront(v[index]))
	}

	this.Unlock()
	return items
}

func (this *Linklist) pushBack(v interface{}) *list.Element {
	return this.L.PushBack(v)
}

func (this *Linklist) PushBack(v interface{}) *list.Element {
	this.Lock()
	e := this.pushBack(v)
	this.Unlock()
	return e
}

func (this *Linklist) BatchPushBack(v []interface{}) []*list.Element {
	length := len(v)

	if length == 0 {
		return []*list.Element{}
	}

	items := make([]*list.Element, 0, length)

	this.Lock()
	for index, _ := range v {
		items = append(items, this.pushBack(v[index]))
	}

	this.Unlock()
	return items
}

func (this *Linklist) Front() interface{} {
	this.RLock()
	value := this.L.Front().Value
	this.RUnlock()
	return value
}

func (this *Linklist) Back() interface{} {
	this.RLock()
	value := this.L.Back().Value
	this.RUnlock()
	return value
}

func (this *Linklist) removeFrontByNum(count int) []interface{} {
	items := make([]interface{}, 0, count)

	for i := 0; i < count; i++ {
		e := this.L.Front()
		if e != nil {
			items = append(items, e.Value)
		}

		this.L.Remove(e)
	}

	return items
}

func (this *Linklist) removeBackByNum(count int) []interface{} {
	items := make([]interface{}, 0, count)

	for i := 0; i < count; i++ {
		e := this.L.Back()
		if e != nil {
			items = append(items, e.Value)
		}

		this.L.Remove(e)
	}

	return items
}

func (this *Linklist) BatchPopFront(count int) []interface{} {
	if count <= 0 {
		return []interface{}{}
	}

	this.Lock()

	total := this.len()

	if total == 0 {
		this.Unlock()
		return []interface{}{}
	}

	if total < count {
		count = total
	}

	items := this.removeFrontByNum(count)

	this.Unlock()
	return items
}

func (this *Linklist) BatchPopBack(count int) []interface{} {
	if count <= 0 {
		return []interface{}{}
	}

	this.Lock()

	total := this.len()

	if total == 0 {
		this.Unlock()
		return []interface{}{}
	}

	if total < count {
		count = total
	}

	items := this.removeBackByNum(count)

	this.Unlock()
	return items
}

func (this *Linklist) PopFront() interface{} {
	if this.Len() < 1 {
		return nil
	}

	this.Lock()

	e := this.L.Front()
	item := this.L.Remove(e)

	this.Unlock()
	return item
}

func (this *Linklist) PopBack() interface{} {
	if this.Len() < 1 {
		return nil
	}

	this.Lock()

	e := this.L.Back()
	item := this.L.Remove(e)

	this.Unlock()
	return item
}

func (this *Linklist) RemainFrontByNum(count int) {
	if count <= 0 {
		return
	}

	this.Lock()

	total := this.len()

	if total == 0 || total <= count {
		this.Unlock()
		return
	}

	delete_size := total - count

	this.removeBackByNum(delete_size)
	this.Unlock()
}

func (this *Linklist) RemainBackByNum(count int) {
	if count <= 0 {
		return
	}

	this.Lock()

	total := this.len()

	if total == 0 || total <= count {
		this.Unlock()
		return
	}

	delete_size := total - count

	this.removeFrontByNum(delete_size)
	this.Unlock()
}
