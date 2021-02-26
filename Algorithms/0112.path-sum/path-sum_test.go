package problem0112

import (
	"testing"

	"github.com/cage1016/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root      *TreeNode
	targetSum int
	ans       bool
}{
	{
		root:      kit.BuildTreeHelper(5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1),
		targetSum: 22,
		ans:       true,
	},
	// 可以有多个 testcase
}

func Test_hasPathSum(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, hasPathSum(tc.root, tc.targetSum), "输入:%v", tc)
	}
}

func Benchmark_hasPathSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			hasPathSum(tc.root, tc.targetSum)
		}
	}
}
