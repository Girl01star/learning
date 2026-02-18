package lru

import "container/list"

type LruCache interface {
	Put(key, value string)
	Get(key string) (string, bool)
}

type entry struct {
	key   string
	value string
}

type lruCache struct {
	capacity int
	ll       *list.List
	items    map[string]*list.Element
}

func NewLruCache(capacity int) LruCache {
	if capacity < 0 {
		capacity = 0
	}
	return &lruCache{
		capacity: capacity,
		ll:       list.New(),
		items:    make(map[string]*list.Element),
	}
}

func (c *lruCache) Get(key string) (string, bool) {
	el, ok := c.items[key]
	if !ok {
		return "", false
	}

	c.ll.MoveToFront(el)
	ent := el.Value.(*entry)
	return ent.value, true
}

func (c *lruCache) Put(key, value string) {
	if c.capacity == 0 {
		return
	}

	if el, ok := c.items[key]; ok {
		ent := el.Value.(*entry)
		ent.value = value
		c.ll.MoveToFront(el)
		return
	}

	if c.ll.Len() >= c.capacity {
		back := c.ll.Back()
		if back != nil {
			ent := back.Value.(*entry)
			delete(c.items, ent.key)
			c.ll.Remove(back)
		}
	}

	el := c.ll.PushFront(&entry{key: key, value: value})
	c.items[key] = el
}
