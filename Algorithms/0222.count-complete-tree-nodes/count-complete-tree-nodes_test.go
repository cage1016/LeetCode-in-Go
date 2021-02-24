package problem0222

import (
	"testing"

	"github.com/cage1016/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root *TreeNode
	ans  int
}{

	{
		root: kit.BuildTreeHelper(1, 2, 3, 4, 5, 6),
		ans:  6,
	},

	// 可以有多个 testcase
}

func Test_countNodes(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, countNodes(tc.root), "输入:%v", tc)
	}
}

func Benchmark_countNodes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			countNodes(tc.root)
		}
	}
}
