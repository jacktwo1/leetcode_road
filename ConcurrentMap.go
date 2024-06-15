package main

import (
	"errors"
	"sync"
	"time"
)

type sp interface {
	Put(key int, val int)                            // 存入key/val，如果该 key 读取的 goroutine 挂起，则唤醒。此方法不会阻塞，时刻都可以立即执行并返回
	Get(key int, timeout time.Duration) (int, error) // 读取一个key，如果 key 不存在阻塞，等待 key 存在或者超时
}

// Map ...
type ConcurrentMap struct {
	sync.Mutex
	mp      map[int]int
	keyToCh map[int]chan struct{} // 判断是否有相应的读协程
}

func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		mp:      make(map[int]int),
		keyToCh: make(map[int]chan struct{}),
	}
}

func (c *ConcurrentMap) Put(key, val int) {
	c.Lock()
	defer c.Unlock()
	c.mp[key] = val
	ch, ok := c.keyToCh[key]
	if !ok {
		return
	}
	select { // 确保多个关闭动作，只会被执行一次，去唤醒所有的读 goroutine
	case <-ch:
		return
	default:
		close(ch)
	}
}

func (c *ConcurrentMap) Get(key int, timeout time.Duration) (int, error) {
	c.Lock()
	v, ok := c.mp[key]
	if ok {
		c.Unlock()
		return v, nil
	}
	ch, ok := c.keyToCh[key]
	if !ok {
		ch = make(chan struct{})
		c.keyToCh[key] = ch
	}
	c.Unlock()
	select {
	case <-time.After(timeout):
		return -1, errors.New("读取超时")
	case <-ch:
	}
	c.Lock()
	v = c.mp[key]
	c.Unlock()
	return v, nil

}
