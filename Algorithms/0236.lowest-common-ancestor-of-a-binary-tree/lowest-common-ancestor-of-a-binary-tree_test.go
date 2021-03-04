package problem0236

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root
 p
 q *TreeNode
	ans *TreeNode
}{



	// 可以有多个 testcase
}

func Test_lowestCommonAncestor(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, lowestCommonAncestor(tc.root, tc.p, tc.q), "输入:%v", tc)
	}
}

func Benchmark_lowestCommonAncestor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			lowestCommonAncestor(tc.root, tc.p, tc.q)
		}
	}
}
