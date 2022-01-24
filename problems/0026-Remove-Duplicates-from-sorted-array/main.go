package main

import "fmt"

func main() {
	nums := []int{1, 1, 2}

	// 解法一
	fmt.Println(removeDuplicates1(nums))

	nums1 := []int{1, 1, 2}
	// 解法二
	fmt.Println(removeDuplicates2(nums1))
}

// 解法一：枚举+哈希表解法
func removeDuplicates1(nums []int) int {
	var count int
	record := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := record[nums[i]]; ok {
			continue
		}
		record[nums[i]] = i
		nums[count] = nums[i]
		count++
	}
	return count
}

// 解法二：通用比较解法(利用有序数组特性)
func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 默认第一个元素不需要处理，所以 count 的起点就是1了
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}

// 解法三：双指针解法
func removeDuplicates3(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	return 0
}
