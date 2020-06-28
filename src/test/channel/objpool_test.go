package channel

import "fmt"

type Item interface {
	run()
}
type Pool interface {
	getItem() *Item
	putItem(*Item)
}

type MyItem struct {
}
type MyPool struct {
	pool chan *Item
}

func (m *MyItem) run() {
	fmt.Println("run myitem")
}

func (m *MyPool) getItem() *Item {
	var item *Item
	item = <-m.pool
	return item
}

func (m *MyPool) putItem(t *Item) {
	m.pool <- t
}
