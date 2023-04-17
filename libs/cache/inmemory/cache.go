package inmemory

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
	"time"
)

type cacheItem struct {
	key   string
	value interface{}
	ttl   *time.Time
}

type ICache interface {
	Set(key string, value interface{}) bool
	Get(key string) (interface{}, error)
}

type cache struct {
	capacity int
	data     map[string]*list.Element
	queue    *list.List
	ticker   *time.Ticker
	ttl      time.Duration
	mu       *sync.Mutex
}

// capacity - вытеснение lru, ttl - время жизни кэша по длительности
func New(capacity int, ttl time.Duration) ICache {
	c := &cache{
		capacity: capacity,
		data:     make(map[string]*list.Element),
		queue:    list.New(),
		ticker:   time.NewTicker(time.Second * 10),
		ttl:      ttl,
		mu:       &sync.Mutex{},
	}

	go c.clean()

	return c
}

func (c *cache) Set(key string, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if val, ok := c.data[key]; ok {
		c.queue.MoveToFront(val)
		val.Value.(*cacheItem).value = value
		return true
	}

	if c.queue.Len() >= c.capacity {
		c.deleteOldest()
	}

	ttl := time.Now().Add(c.ttl)

	item := &cacheItem{
		key:   key,
		value: value,
		ttl:   &ttl,
	}

	el := c.queue.PushFront(item)
	c.data[item.key] = el

	return true
}

func (c *cache) Get(key string) (interface{}, error) {
	op := "cache.Get"

	c.mu.Lock()
	defer c.mu.Unlock()

	if el, ok := c.data[key]; ok {
		c.queue.MoveToFront(el)
		return el.Value.(*cacheItem).value, nil
	}

	return nil, fmt.Errorf("%s: %w", op, errors.New("cannot retrieve cache"))
}

func (c *cache) deleteOldest() {
	if el := c.queue.Back(); el != nil {
		c.deleteItem(el)
	}
}

func (c *cache) clean() {
	for {
		<-c.ticker.C
		c.mu.Lock()
		for _, val := range c.data {
			el := val.Value.(*cacheItem)
			if el.ttl != nil {
				if time.Now().After(*el.ttl) {
					c.deleteItem(val)
				}
			}
		}
		c.mu.Unlock()
	}
}

func (c *cache) deleteItem(el *list.Element) {
	item := c.queue.Remove(el).(*cacheItem)
	delete(c.data, item.key)
}
