package variables

import (
	"testing"
)

var (
	i int = 1
)

func TestVariable(t *testing.T) {
	t.Log(i)
	i++
}

func TestVariable2(t *testing.T) {
	t.Log(i)
}
