/*
给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

请必须使用时间复杂度为 O(log n) 的算法。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/search-insert-position
*/
package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	target := 4

	fmt.Println(searchInsert(nums, target))
}

// 解法一：普通比较解法
func searchInsert(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}
