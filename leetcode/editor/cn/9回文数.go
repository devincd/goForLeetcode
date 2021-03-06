//给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
//
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
//
//
// 例如，121 是回文，而 123 不是。
//
//
//
//
// 示例 1：
//
//
//输入：x = 121
//输出：true
//
//
// 示例 2：
//
//
//输入：x = -121
//输出：false
//解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
//
//
// 示例 3：
//
//
//输入：x = 10
//输出：false
//解释：从右向左读, 为 01 。因此它不是一个回文数。
//
//
//
//
// 提示：
//
//
// -2³¹ <= x <= 2³¹ - 1
//
//
//
//
// 进阶：你能不将整数转为字符串来解决这个问题吗？
//
// Related Topics 数学 👍 2114 👎 0
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("121 -> ", Reverse(121))
	fmt.Println("12456 -> ", Reverse(12456))
	fmt.Println("-321 -> ", Reverse(-321))
}

//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(x int) bool {
	// 时间复杂度O(n)
	stringX := strconv.Itoa(x)
	for i := 0; i < len(stringX)/2; i++ {
		if stringX[i] != stringX[len(stringX)-1-i] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

// Reverse 数字反转
func Reverse(x int) int {
	var cur int
	for x != 0 {
		// 1.取数字最后一位数
		tail := x % 10
		// 2.拼接反转后的新数
		cur = cur*10 + tail
		// 3.去掉最后一位数
		x = x / 10
	}
	return cur
}
