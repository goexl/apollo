package apollo_test

import (
	"testing"

	"github.com/goexl/apollo"
)

func TestNew(t *testing.T) {
	client := apollo.New().Build()
	if nil == client {
		t.Error("创建客户端出错")
	}
}
