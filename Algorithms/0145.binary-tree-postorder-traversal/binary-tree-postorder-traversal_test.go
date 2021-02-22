package problem0145

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cage1016/LeetCode-in-Go/kit"
)

// tcs is testcase slice
var tcs = []struct {
	root *TreeNode
	ans  []int
}{
	{
		root: kit.BuildTreeHelper(5, 4, 6, 1, 2),
		ans:  []int{1, 2, 4, 6, 5},
	},

	// 可以有多个 testcase
}

func Test_postorderTraversal(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, postorderTraversal(tc.root), "输入:%v", tc)
	}
}

func Benchmark_postorderTraversal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			postorderTraversal(tc.root)
		}
	}
}
