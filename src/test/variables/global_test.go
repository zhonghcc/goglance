package variables

import (
	v1 "goglance/src/test/variables/pkg1"
	v2 "goglance/src/test/variables/pkg2"
	"testing"
)

func TestGolbal(t *testing.T) {
	t.Log("nn")

	t.Log(v2.Pkg2Int)
	v2.Pkg2Int = v2.Pkg2Int + 1
	t.Log(v2.Pkg2Int)
	v1.Add()
	t.Log(v2.Pkg2Int)
}
