package util

import (
	"sync"
)

// RingQueue 固定数量循环队列
type RingQueue struct {
	data       []interface{}
	head, size int
	capacity   int
	mu         sync.RWMutex
}

func NewRingQueue(capacity int) *RingQueue {
	return &RingQueue{
		data:     make([]interface{}, capacity),
		capacity: capacity,
	}
}

// Data 原始数据（会被篡改）
func (that *RingQueue) Data() []interface{} {
	return that.data
}

func (that *RingQueue) Size() int {
	return that.size
}

func (that *RingQueue) Capacity() int {
	return that.capacity
}

// Enqueue 入队，满了会覆盖最旧的
func (that *RingQueue) Enqueue(val interface{}) {
	that.mu.Lock()
	defer that.mu.Unlock()

	if that.size < that.capacity {
		that.data[(that.head+that.size)%that.capacity] = val
		that.size++
	} else {
		// 队列满了：覆盖最旧元素
		that.data[that.head] = val
		that.head = (that.head + 1) % that.capacity
	}
}

// Dequeue 出队
func (that *RingQueue) Dequeue() (interface{}, bool) {
	that.mu.Lock()
	defer that.mu.Unlock()

	if that.size == 0 {
		return nil, false
	}
	val := that.data[that.head]
	that.head = (that.head + 1) % that.capacity
	that.size--
	return val, true
}

// Values 获取队列内容（从旧到新）
func (that *RingQueue) Values() []interface{} {
	that.mu.RLock()
	defer that.mu.RUnlock()

	values := make([]interface{}, that.size)
	for i := 0; i < that.size; i++ {
		values[i] = that.data[(that.head+i)%that.capacity]
	}
	return values
}

// View 返回一个只读视图
// f func(i int, v interface{}) bool 返回 false 则不再继续遍历提供视图，view 找到需要的参数后应该 return false 结束无意义的遍历
func (that *RingQueue) View(f func(index int, value interface{}) bool) {
	that.mu.RLock()
	defer that.mu.RUnlock()

	for i := 0; i < that.size; i++ {
		if !f(i, that.data[(that.head+i)%that.capacity]) {
			break
		}
	}
}

// Latest  最新数据
func (that *RingQueue) Latest() (interface{}, bool) {
	that.mu.RLock()
	defer that.mu.RUnlock()

	if that.size == 0 {
		return nil, false
	}
	latestIdx := (that.head + that.size - 1) % that.capacity
	return that.data[latestIdx], true
}

// Oldest 获取最旧的元素
func (that *RingQueue) Oldest() (interface{}, bool) {
	that.mu.RLock()
	defer that.mu.RUnlock()

	if that.size == 0 {
		return nil, false
	}
	return that.data[that.head], true
}

// Get 获取指定位置的元素（0 = 最旧）
func (that *RingQueue) Get(index int) (interface{}, bool) {
	that.mu.RLock()
	defer that.mu.RUnlock()

	if index < 0 || index >= that.size {
		return nil, false
	}
	pos := (that.head + index) % that.capacity
	return that.data[pos], true
}
