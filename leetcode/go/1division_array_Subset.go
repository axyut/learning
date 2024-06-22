package main

import (
	"fmt"
)

func solution(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// Initialize dp array where dp[i] will store the length of the largest subset ending with nums[i]
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1 // Initialize each element as a subset of size 1 (itself)
	}

	// Traverse each pair (i, j) such that j < i
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if (nums[i]+nums[j])%target != 0 {
				if dp[i] < dp[j]+1 {
					dp[i] = dp[j] + 1
				}
			}
		}
	}

	// Find the maximum value in dp array
	maxSubsetSize := 0
	for _, size := range dp {
		if size > maxSubsetSize {
			maxSubsetSize = size
		}
	}

	return maxSubsetSize
}

func DivArraySubset_1() {
	fmt.Println(solution([]int{1, 7, 2, 4}, 3)) // Output: 3
	fmt.Println(solution([]int{1}, 3))          // Output: 1
}

// Division and Array Subsets
// Explanation
// In the solution function, a different number of integer array nums and integer value target were given as arguments. Please write a program to return the maximum value of the subset as the return value of the solution function in integer that satisfies the conditions where the elements of nums array are combined and the sum of any two numbers is not divisible by target. However, a subset consisting of a single element must satisfy the above condition.

// Examples
// Example 1
// Input: nums = [1,7,2,4], target = 3
// Output: 3
// Explanation: Since the subset [1,7,4]satisfies the condition where none of   1+7 7+4 4+1 are divisible by 3 and as there is no larger subset than itself、the answer is 3, which is the array length of [1,7,4]
// Example 2:
// Input: nums = [1], target = 3
// Output: 1
// Explanation:Since the only subset [1] cannot form a combination of two numbers, it automatically satisfies the condition that the sum of any two numbers is not divisible by target thus the answer is 1 ,which is the array length of the subset.
// Constraints
// 1 ≦ len(nums), target ≦ 10000
// nums[i] is any int satisfying 1 ≦ nums[i] < inf
// Each element of nums is different int.
