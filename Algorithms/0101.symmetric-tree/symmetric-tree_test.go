package problem0101

import (
	"testing"

	"github.com/cage1016/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root *TreeNode
	ans  bool
}{

	{
		root: kit.BuildTreeHelper(1, 2, 2, 3, 4, 4, 3),
		ans:  true,
	},

	{
		root: kit.BuildTreeHelper(1, 2, 2, nil, 3, nil, 3),
		ans:  false,
	},

	// 可以有多个 testcase
}

func Test_isSymmetric(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, isSymmetric(tc.root), "输入:%v", tc)
	}
}

func Benchmark_isSymmetric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isSymmetric(tc.root)
		}
	}
}
