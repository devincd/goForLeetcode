package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	target := 0

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

// 解法二：二分查找解法
func searchInsert1(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}
