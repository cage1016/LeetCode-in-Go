package problem0108

import (
	"testing"

	"github.com/cage1016/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  *TreeNode
}{
	{
		nums: []int{-10, -3, 0, 5, 9},
		ans:  kit.BuildTreeHelper(0, -3, 9, -10, nil, 5),
	},

	// 可以有多个 testcase
}

func Test_sortedArrayToBST(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, sortedArrayToBST(tc.nums), "输入:%v", tc)
	}
}

func Benchmark_sortedArrayToBST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			sortedArrayToBST(tc.nums)
		}
	}
}
