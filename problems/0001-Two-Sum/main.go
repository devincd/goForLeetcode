// 0001. Two Sum
// https://leetcode-cn.com/problems/two-sum/
package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	// 解法一
	fmt.Println(twoSum1(nums, target))

	// 解法二
	fmt.Println(twoSum2(nums, target))
}

// 解法一：暴力枚举
// 时间复杂度：O(n*n)
// 空间复杂度：O(1)
func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 解法二：哈希表
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func twoSum2(nums []int, target int) []int {
	record := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if index, ok := record[target-nums[i]]; ok {
			return []int{index, i}
		}

		record[nums[i]] = i
	}

	return nil
}
