package problem0257

import (
	"testing"

	"github.com/cage1016/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root *TreeNode
	ans  []string
}{

	{
		root: kit.BuildTreeHelper(1, 2, 3, nil, 5),
		ans:  []string{"1->2->5", "1->3"},
	},

	// 可以有多个 testcase
}

func Test_binaryTreePaths(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, binaryTreePaths(tc.root), "输入:%v", tc)
	}
}

func Benchmark_binaryTreePaths(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			binaryTreePaths(tc.root)
		}
	}
}
