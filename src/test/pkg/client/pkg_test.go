package client

import (
	"testing"

	//相对路径，不建议
	server "../server"
)

func TestPkg(t *testing.T) {
	t.Log(server.GetIntSqure(5))
}
