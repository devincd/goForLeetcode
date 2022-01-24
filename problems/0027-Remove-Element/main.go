package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3

	fmt.Println(removeElement1(nums, val))
}

// 解法一：普通比较解法
func removeElement1(nums []int, val int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}
		nums[count] = nums[i]
		count++
	}

	return count
}
