package main

func moveZeroes(nums []int) {
	last, first := 0, 0
	for _, v := range nums {
		if v == 0 {
			last++
		} else {
			nums[first] = v
			first++
		}
	}
	for i := 0; i < last; i++ {
		nums[first+i] = 0
	}
	return
}
