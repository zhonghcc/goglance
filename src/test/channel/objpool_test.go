package channel

import (
	"fmt"
	"testing"
)

type Item interface {
	run()
	init()
}
type Pool interface {
	getItem() Item
	putItem(Item)
	init(size int)
	getPoolSize() int
}

type MyItem struct {
	no int
}
type MyPool struct {
	pool chan Item
}

func (m *MyItem) run() {
	fmt.Printf("run myitem %p,%d\n", m, m.no)
}

var mark int = 1

func (m *MyItem) init() {
	m.no = mark
	mark++
	fmt.Printf("init myitem %p,%d\n", m, m.no)
}

func (m *MyPool) getItem() Item {
	var item Item
	item = <-m.pool
	return item
}

func (m *MyPool) putItem(t Item) {
	m.pool <- t
}

func (m *MyPool) getPoolSize() int {
	return len(m.pool)
}

func (m *MyPool) init(size int) {
	m.pool = make(chan Item, size)
	fmt.Println("init pool start")
	for i := 0; i < size; i++ {
		item := new(MyItem)
		item.init()
		m.pool <- item
	}
	fmt.Println("size of pool ", m.getPoolSize())
}

func TestPool(t *testing.T) {

	var pool Pool
	pool = new(MyPool)
	pool.init(10)

	for i := 0; i < 20; i++ {
		var item Item = pool.getItem()
		item.run()
		pool.putItem(item)
	}
}
