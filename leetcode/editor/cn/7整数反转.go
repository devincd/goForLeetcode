//给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
//
// 如果反转后整数超过 32 位的有符号整数的范围 [−2³¹, 231 − 1] ，就返回 0。
//假设环境不允许存储 64 位整数（有符号或无符号）。
//
//
//
// 示例 1：
//
//
//输入：x = 123
//输出：321
//
//
// 示例 2：
//
//
//输入：x = -123
//输出：-321
//
//
// 示例 3：
//
//
//输入：x = 120
//输出：21
//
//
// 示例 4：
//
//
//输入：x = 0
//输出：0
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
// Related Topics 数学 👍 3569 👎 0
package main

import "math"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func reverse(x int) int {
	var cur int
	for x != 0 {
		// 1.取数字的最后一位
		tail := x % 10
		// 2.拼接反转后的数字
		cur = cur*10 + tail
		// 3.去掉数字的最后一位
		x = x / 10
	}
	// 判断是否越界
	if cur > math.MaxInt32 || cur < math.MinInt32 {
		return 0
	}
	return cur
}

//leetcode submit region end(Prohibit modification and deletion)
