package problem0637

import (
	"testing"

	"github.com/cage1016/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root *TreeNode
	ans  []float64
}{
	{
		root: kit.BuildTreeHelper(3, 9, 20, 15, 7),
		ans:  []float64{3.00000, 14.50000, 11.00000},
	},

	// 可以有多个 testcase
}

func Test_averageOfLevels(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, averageOfLevels(tc.root), "输入:%v", tc)
	}
}

func Benchmark_averageOfLevels(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			averageOfLevels(tc.root)
		}
	}
}
