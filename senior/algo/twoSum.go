package main

// 两数求和
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if target == nums[i]+nums[j] {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// hash法
// leetcode的测试用例只需要3ms，js相同算法需要64ms，这就是选型的区别
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		// j - map中存储的数组索引(value), ok - 判断complement是否存在于map中(key)
		if j, ok := m[complement]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return []int{}
}
