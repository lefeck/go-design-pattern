package main

import "fmt"

type EvictionAlgo interface {
	Evict(c *Cache)
}

type LRU struct {
}

func (l *LRU) Evict(c *Cache) {
	fmt.Println("Evicting by lru strtegy")
}

type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func NewCache(e EvictionAlgo) *Cache {
	return &Cache{
		storage:      make(map[string]string),
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) evict() {
	c.evictionAlgo.Evict(c)
	c.capacity--
}

func (c *Cache) Add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) Get(key string) {
	delete(c.storage, key)
}
