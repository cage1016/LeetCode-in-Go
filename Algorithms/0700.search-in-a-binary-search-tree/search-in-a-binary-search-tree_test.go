package problem0700

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root *TreeNode
 val int
	ans *TreeNode
}{



	// 可以有多个 testcase
}

func Test_searchBST(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, searchBST(tc.root, tc.val), "输入:%v", tc)
	}
}

func Benchmark_searchBST(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			searchBST(tc.root, tc.val)
		}
	}
}
