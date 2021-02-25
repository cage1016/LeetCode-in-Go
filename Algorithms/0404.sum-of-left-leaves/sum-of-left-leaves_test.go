package problem0404

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
		root: kit.BuildTreeHelper(3, 9, 20, nil, nil, 15, 7),
		ans:  24,
	},

	// 可以有多个 testcase
}

func Test_sumOfLeftLeaves(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, sumOfLeftLeaves(tc.root), "输入:%v", tc)
	}
}

func Benchmark_sumOfLeftLeaves(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			sumOfLeftLeaves(tc.root)
		}
	}
}
