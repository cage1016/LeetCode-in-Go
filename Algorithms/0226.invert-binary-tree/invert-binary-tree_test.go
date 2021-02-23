package problem0226

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root *TreeNode
	ans *TreeNode
}{



	// 可以有多个 testcase
}

func Test_invertTree(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, invertTree(tc.root), "输入:%v", tc)
	}
}

func Benchmark_invertTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			invertTree(tc.root)
		}
	}
}
