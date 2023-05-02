package hw04lrucache

type Key string

type Item struct {
	key   Key
	value interface{}
}

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (l *lruCache) purge() {
	if el := l.queue.Back(); el != nil {
		delete(l.items, el.Value.(*Item).key)
		l.queue.Remove(el)
		el = nil
	}
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	if l.capacity == 0 {
		return false
	}
	if el, exists := l.items[key]; exists {
		l.queue.MoveToFront(el)
		el.Value.(*Item).value = value
		return true
	}

	if l.queue.Len() == l.capacity {
		l.purge()
	}

	item := &Item{
		key:   key,
		value: value,
	}

	el := l.queue.PushFront(item)
	l.items[key] = el

	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	element, exists := l.items[key]
	if !exists {
		return nil, false
	}
	l.queue.MoveToFront(element)
	return element.Value.(*Item).value, true
}

func (l *lruCache) Clear() {
	l.items = nil
	l.queue = nil
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
