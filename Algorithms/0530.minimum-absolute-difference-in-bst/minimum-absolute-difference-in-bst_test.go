package problem0530

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
		root: kit.BuildTreeHelper(1, nil, 5, 3),
		ans:  2,
	},

	// 可以有多个 testcase
}

func Test_getMinimumDifference(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, getMinimumDifference(tc.root), "输入:%v", tc)
	}
}

func Benchmark_getMinimumDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			getMinimumDifference(tc.root)
		}
	}
}
