package problem0501

import (
	"testing"

	"github.com/cage1016/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root *TreeNode
	ans  []int
}{
	{
		root: kit.BuildTreeHelper(1, nil, 2, 2),
		ans:  []int{2},
	},
	{
		root: kit.BuildTreeHelper(1, nil, 2),
		ans:  []int{1, 2},
	},
	// 可以有多个 testcase
}

func Test_findMode(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, findMode(tc.root), "输入:%v", tc)
	}
}

func Benchmark_findMode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findMode(tc.root)
		}
	}
}
