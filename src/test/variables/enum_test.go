package variables

import (
	"testing"
)

const (
	Readable = 1 << iota
	Writeable
	Executable
)

const (
	Monday = 1 + iota
	Tuesday
	WendesDay
)

func TestEnum(t *testing.T) {
	t.Logf("%04b %04b %04b", Readable, Writeable, Executable)
	t.Logf("%d %d %d", Monday, Tuesday, WendesDay)
}
