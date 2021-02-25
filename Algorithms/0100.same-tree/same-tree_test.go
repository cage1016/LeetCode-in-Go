package problem0100

import (
	"testing"

	"github.com/cage1016/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	p   *TreeNode
	q   *TreeNode
	ans bool
}{

	{
		p:   kit.BuildTreeHelper(1),
		q:   kit.BuildTreeHelper(1, nil, 2),
		ans: false,
	},

	// 可以有多个 testcase
}

func Test_isSameTree(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, isSameTree(tc.p, tc.q), "输入:%v", tc)
	}
}

func Benchmark_isSameTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isSameTree(tc.p, tc.q)
		}
	}
}
