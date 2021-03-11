package problem0669

import (
	"testing"

	"github.com/cage1016/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root *TreeNode
	low  int
	high int
	ans  *TreeNode
}{
	{
		root: kit.BuildTreeHelper(3),
		low:  2,
		high: 2,
		ans:  nil,
	},
	// 可以有多个 testcase
}

func Test_trimBST(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, trimBST(tc.root, tc.low, tc.high), "输入:%v", tc)
	}
}

func Benchmark_trimBST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			trimBST(tc.root, tc.low, tc.high)
		}
	}
}
