package container

import (
	"fmt"
	"testing"
)

type Days int

const (
	Monday  Days = 1
	TuesDay Days = 2
	Sunday  Days = 7
)

func TestFactory(t *testing.T) {
	factory := map[Days]func(){}

	factory[Monday] = weekday
	factory[TuesDay] = weekday
	factory[Sunday] = weekend

	doFunc(factory, Monday)
	doFunc(factory, Sunday)
}

func weekend() {
	fmt.Println("today is weekend")
}
func weekday() {
	fmt.Println("today is weekday")
}

func doFunc(m map[Days]func(), day Days) {
	if value, exist := m[day]; exist == false {
		fmt.Errorf("not exist key %d", day)
	} else {
		value()
	}
}
