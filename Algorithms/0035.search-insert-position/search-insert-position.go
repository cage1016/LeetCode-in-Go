/*
* @lc app=leetcode id=35 lang=golang
 */

package problem0035

// @lc code=start
func searchInsert(nums []int, target int) int {
	// method 1, O(n)
	// for i := 0; i < len(nums); i++ {
	// 	if nums[i] >= target {
	// 		return i
	// 	}
	// }
	// return len(nums)

	// method 2, O(logN), [left, right]
	left := 0
	right := len(nums) - 1
	for left <= right {
		middle := left + ((right - left) / 2)
		if nums[middle] > target {
			right = middle - 1
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			return middle
		}
	}
	return right + 1

	// // method 3, O(logN), [left, right)
	// left := 0
	// right := len(nums)
	// for left < right {
	// 	middle := left + ((right - left) / 2)
	// 	if nums[middle] > target {
	// 		right = middle
	// 	} else if nums[middle] < target {
	// 		left = middle + 1
	// 	} else {
	// 		return middle
	// 	}
	// }
	// return right
}

// @lc code=end
