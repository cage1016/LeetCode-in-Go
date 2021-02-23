package problem0116

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root *Node
	ans *Node
}{



	// 可以有多个 testcase
}

func Test_connect(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, connect(tc.root), "输入:%v", tc)
	}
}

func Benchmark_connect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			connect(tc.root)
		}
	}
}
