/*
* @lc app=leetcode id=1576 lang=golang
 */

package problem1576

// @lc code=start
func modifyString(s string) string {
	res := []byte(s)
	for i := 0; i < len(s); i++ {
		if s[i] == '?' {
			for j := byte('a'); j <= 'z'; j++ {
				if (i == 0 || res[i-1] != j) && (i == len(s)-1 || s[i+1] != j) {
					res[i] = j
					break
				}
			}
		}
	}
	return string(res)
}

// @lc code=end
