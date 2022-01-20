## LeetCode第1号问题：两数之和

### 题目描述
给定一个整数数组`nums`和一个整数目标值`target`，请你在该数组中找出和为目标值`target`的那两个整数
，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按照任意顺序返回答案。

#### 示例
示例1：
```
输入: nums = [2,7,11,15], target = 9
输出：[0,1]
解释: 因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
```
示例2：
```
输入: nums = [3,2,4], target = 6
输出：[1,2]
```

示例3：
```
输入: nums = [3,3], target = 6
输出：[0,1]
```
### 题目解析
#### 解法一：暴力枚举
思路及算法：

最容易想到的方法是枚举数组中的每一个数`x`，寻找数组中是否存在`target - x`。当我们使用遍历整个数组的方式寻找
`target - x`时，需要注意到每一个位于`x`之前的元素都已经和`x`匹配过，因此不需要再进行匹配。
而每一个元素不能被使用两次，所以我们只需要在`x`后面的元素中寻找`target - x`。

复杂度分析：
- 时间复杂度：O(n²)，其中`n`是数组中的元素数量。最坏情况下数组中任意两个数都要被匹配一次。
- 空间复杂度：O(1)。

#### 解法二：哈希表
思路及算法：

注意到解法一的时间复杂度较高的原因是寻找`target - x`的时间复杂度过高。因此，我们需要一种更优秀的方法，
能够快速寻找数组中是否存在目标元素。如果存在，我们需要找出它的索引。
使用哈希表，可以将寻找`target - x`的时间复杂度降低，从O(n)降低到O(1)。

这样我们创建一个哈希表，对于每一个`x`，我们首先查询哈希表中是否存在`target - x`，然后将`x`插入到哈希表中，
即保证不会让`x`和自己匹配。

复杂度分析：
- 时间复杂度：O(n)，其中`n`是数组中元素数量。对于每一个元素`x`，我们可以O(1)地寻找`target - x`。
- 空间复杂度：O(n)，其中`n`是数组中的元素数量。主要是哈希表的开销。

### 动画描述
TODO

### 代码实现
解法一：暴力枚举
```go
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
```

解法二：哈希表
```go
func twoSum(nums []int, target int) []int {
	// 初始化哈希表，用于存在 nums 中的 index 和 value。
	// 格式为 record[value] = index
	record := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		// 查询哈希表中是否存在 target - value，如果存在，返回值即为需要寻找的下标。
		if index, ok := record[target-nums[i]]; ok {
			return []int{index, i}
		}
		// 哈希表中没有找到 target - value，以 record[value] = index 格式存储到哈希表中。
		record[nums[i]] = i
	}

	return nil
}
```