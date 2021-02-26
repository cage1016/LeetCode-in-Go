package problem0513

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

	// {
	// 	root: kit.BuildTreeHelper(1, 2, 3, 4, nil, 5, 6, nil, nil, 7),
	// 	ans:  7,
	// },
	{
		root: kit.BuildTreeHelper(1),
		ans:  1,
	},

	// 可以有多个 testcase
}

func Test_findBottomLeftValue(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, findBottomLeftValue(tc.root), "输入:%v", tc)
	}
}

func Benchmark_findBottomLeftValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			findBottomLeftValue(tc.root)
		}
	}
}
