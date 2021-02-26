package problem0113

import (
	"testing"

	"github.com/cage1016/LeetCode-in-Go/kit"
	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	root      *TreeNode
	targetSum int
	ans       [][]int
}{
	{
		root:      kit.BuildTreeHelper(-260, -202, -903, -980, -570, -858, 218, 764, -300, 205, nil, -35, nil, nil, -204, 950, -769, 258, -652, 614, -584, 76, 817, -192, nil, nil, -114, 880, nil, -200, 71, 671, 344, 801, 193, -18, 876, -920, -730, 222, 679, nil, -680, nil, nil, nil, -859, 744, -261, 692, nil, -341, -163, nil, nil, 482, -979, 205, nil, 146, 165, 801, 100, -656, 714, -629, 995, 474, 307, -581, -150, -941, nil, nil, nil, -937, -69, -23, 82, nil, -139, -591, nil, -453, -861, -370, nil, nil, nil, 216, 233, nil, 430, nil, 5, -110, nil, nil, -660, 624, -510, -588, nil, nil, 381, nil, 368, 559, nil, 521, -301, nil, 522, 379, nil, nil, nil, nil, 456, 519, nil, nil, 482, 349, nil, nil, 19, nil, nil, 288, -811, nil, -372, nil, nil, -536, nil, -404, -457, -740, 860, nil, nil, -636, nil, nil, 342, -874, -462, -504, 781, 855, -392, nil, nil, nil, 406, nil, -758, 541, nil, -947, nil, nil, nil, nil, nil, -964, nil, 600, -45, nil, nil, nil, nil, nil, nil, nil, nil, nil, -194, nil, nil, nil, -802, nil, nil, nil, -3, nil, -792, 672, 643, nil, 14, nil, nil, 489, 457, nil, nil, nil, nil, 412, nil, 558, nil, nil, nil, nil, -846, 158, -146, nil, nil, -76, -650, nil, -782, nil, -127, nil, -678, nil, nil, nil, nil, nil, nil, -464, -426, nil, -366, nil, nil, nil, nil, nil, 81, -607, 716, nil, nil, -213, nil, 379, nil, nil, nil, nil, 644, 445, nil, nil, -419, -845, -720, nil, nil, -915, nil, nil, nil, nil, nil, nil, -686, 594, -243, nil, 496, nil, 907, nil, nil, nil, nil, nil, nil, 579, 873, 702, nil, nil, nil, -834, nil, nil, nil, nil, nil, -300, -214, -466, nil, nil, 972, nil, nil, nil, 814, nil, -940, nil, 763, nil, nil, nil, nil, -449, -844, nil, nil, nil, nil, -47),
		targetSum: -243,
		ans:       [][]int{{-260, -903, -858, -35, 817, 222, 307, -301, -947, -76, 445, 579, 814, -47}},
	},
	{
		root:      kit.BuildTreeHelper(-6, nil, -3, -6, 0, -6, -5, 4, nil, nil, nil, 1, 7),
		targetSum: -21,
		ans:       [][]int{{-6, -3, -6, -6}},
	},
	{
		root:      kit.BuildTreeHelper(5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1),
		targetSum: 22,
		ans:       [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
	},
	// 可以有多个 testcase
}

func Test_pathSum(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, pathSum(tc.root, tc.targetSum), "输入:%v", tc)
	}
}

func Benchmark_pathSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			pathSum(tc.root, tc.targetSum)
		}
	}
}
