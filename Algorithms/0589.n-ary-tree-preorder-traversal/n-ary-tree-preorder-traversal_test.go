package problem0589

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root *Node
	ans []int
}{



	// 可以有多个 testcase
}

func Test_preorder(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, preorder(tc.root), "输入:%v", tc)
	}
}

func Benchmark_preorder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			preorder(tc.root)
		}
	}
}
