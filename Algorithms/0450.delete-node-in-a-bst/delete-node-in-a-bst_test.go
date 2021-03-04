package problem0450

import (
	"testing"

	"github.com/cage1016/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root *TreeNode
	key  int
	ans  *TreeNode
}{
	{
		root: kit.BuildTreeHelper(1, nil, 2),
		key:  1,
		ans:  kit.BuildTreeHelper(2),
	},
	{
		root: kit.BuildTreeHelper(5, 3, 6, 2, 4, nil, 7),
		key:  7,
		ans:  kit.BuildTreeHelper(5, 3, 6, 2, 4),
	},
	{
		root: kit.BuildTreeHelper(5, 3, 6, 2, 4, nil, 7),
		key:  5,
		ans:  kit.BuildTreeHelper(6, 3, 7, 2, 4),
	},
	{
		root: kit.BuildTreeHelper(5, 3, 6, 2, 4, nil, 7),
		key:  3,
		ans:  kit.BuildTreeHelper(5, 4, 6, 2, nil, nil, 7),
	},
	{
		root: kit.BuildTreeHelper(0),
		key:  0,
		ans:  nil,
	},
	// 可以有多个 testcase
}

func Test_deleteNode(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, deleteNode(tc.root, tc.key), "输入:%v", tc)
	}
}

func Benchmark_deleteNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			deleteNode(tc.root, tc.key)
		}
	}
}
