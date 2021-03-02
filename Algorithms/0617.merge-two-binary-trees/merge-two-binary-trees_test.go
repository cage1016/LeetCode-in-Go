package problem0617

import (
	"testing"

	"github.com/cage1016/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root1 *TreeNode
	root2 *TreeNode
	ans   *TreeNode
}{
	{
		root1: kit.BuildTreeHelper(1, 2, nil, 3),
		root2: kit.BuildTreeHelper(1, nil, 2, nil, 3),
		ans:   kit.BuildTreeHelper(2, 2, 2, 3, nil, nil, 3),
	},
	{
		root1: nil,
		root2: kit.BuildTreeHelper(1),
		ans:   kit.BuildTreeHelper(1),
	},
	{
		root1: kit.BuildTreeHelper(1, 3, 2, 5),
		root2: kit.BuildTreeHelper(2, 1, 3, nil, 4, nil, 7),
		ans:   kit.BuildTreeHelper(3, 4, 5, 5, 4, nil, 7),
	},

	// 可以有多个 testcase
}

func Test_mergeTrees(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, mergeTrees(tc.root1, tc.root2), "输入:%v", tc)
	}
}

func Benchmark_mergeTrees(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			mergeTrees(tc.root1, tc.root2)
		}
	}
}
