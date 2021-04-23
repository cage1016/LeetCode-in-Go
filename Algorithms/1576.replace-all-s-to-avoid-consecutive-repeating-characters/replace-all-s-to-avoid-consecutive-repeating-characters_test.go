package problem1576

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	ans string
}{
	// {
	// 	"?zs", "azs",
	// },
	{
		"ab?ac?", "abcaca",
	},
	// {
	// 	"???a", "",
	// },
	// 可以有多个 testcase
}

func Test_modifyString(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, modifyString(tc.s), "输入:%v", tc)
	}
}

func Benchmark_modifyString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			modifyString(tc.s)
		}
	}
}
