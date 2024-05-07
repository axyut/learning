package main

import "fmt"

func main() {
	// execute different DSAs
	ret := twoSum([]int{-1, -2, -3, -4, -5}, -8)
	fmt.Println(ret)
	palinTestStr("watch")
	palinTestNum(12321)
	// filter, map, reduce
	a := []int{1, 2, 6, 3, 4, 5, 6, 7, 8, 9, 10, 2}
	onlyEven := filter(a, func(val int) bool {
		return val%2 == 0
	})
	fmt.Println(onlyEven, " even numbers")
	PrimeChecker(12)
	fibo(8)
	// fizzbuzz()
	removeDup(a)
	// remove duplicates
	// tower of hanoi
	// ceaser cipher to string
	// hashmaps
}

func twoSum(nums []int, target int) []int {
	arr := []int{0, 1}
	// found := false
	for i, v := range nums {
		if i+2 > len(nums) {
			break
		}
		for i1 := i + 1; i < len(nums); i++ {
			if nums[i1]+v == target {
				arr[0] = i
				arr[1] = i1
				return arr
			}
		}
		// for i1, v1 := range nums {
		// 	if i1 == i {
		// 		continue
		// 	}
		// 	if v+v1 == target {
		// 		arr[0] = i
		// 		arr[1] = i1
		// 		return arr
		// 	}
		// }
	}
	return arr
}

func removeDup(elements []int) {
	// Use map to record duplicates as we find them.
	encountered := map[int]bool{}
	result := []int{}

	for v := range elements {
		if encountered[elements[v]] {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	fmt.Println(result)
}

// func removeDup(arr []int) {
// 	new := make([]int, len(arr))

// 	for i1, v1 := range arr {
// 		fmt.Println(v1, "top")
// 		for _, v2 := range new {
// 			fmt.Println(v1, "deep")
// 			if v1 == v2 {
// 				fmt.Println(v1, "here")
// 				break
// 			} else {
// 				fmt.Println(v1)
// 				new[i1] = v1
// 				fmt.Println(new)
// 				continue
// 			}
// 		}
// 	}

// 	fmt.Println(new)
// }

// fizz at multiplies of 3, buzz at multiples of 5 and fizzbuzz at no. that both multiplies
func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Printf("%d. FizzBuzz\n", i)
				continue
			}
			fmt.Printf("%d. Fizz\n", i)
			continue
		} else {
			fmt.Printf("%d. Buzz\n", i)
		}
	}
}

// 0 1 1 2 3 5 8 12
func fibo(nth int) int {
	nextTerm, t1, t2 := 0, 0, 1
	for i := 1; i <= nth; i++ {
		if i == 1 || i == 2 {
			continue
		}
		nextTerm = t1 + t2
		t1 = t2
		t2 = nextTerm
	}
	fmt.Printf("%d. term in fibonacci is %d\n", nth, t2)
	return t2
}

func PrimeChecker(num int) bool {
	div := 2

	for num > div {
		if num%div == 0 {
			fmt.Printf("%d is not a Prime.\n", num)
			return false
		} else {
			div++
		}
	}
	fmt.Printf("%d is a Prime.\n", num)
	return true

}

func palinTestStr(input_text string) bool {
	leng := len(input_text)

	var last string
	for i := leng; i >= 1; i-- {
		last += string(input_text[i-1])
	}

	if input_text == last {
		fmt.Println(input_text + " is Palindrome.")
		return true
	}
	fmt.Println(input_text + " is not Palindrome.")
	return false
}

// func palinTestStr(input string) bool {
// 	for i := 0; i < len(input)/2; i++ {
// 		lastChar := len(input) - 1
// 		if input[i] != string[lastChar-i] {
// 			fmt.Println(input + " is not Palindrome.")
// 			return false
// 		}
// 	}
// 	fmt.Println(input + " is Palindrome.")
// 	return true
// }

func palinTestNum(input int) bool {
	var temp, last, ulto int
	temp = input
	for temp > 0 {
		last = temp % 10
		ulto = (ulto * 10) + last
		temp = temp / 10
	}
	if input == ulto {
		fmt.Printf("%d is Palindrome.\n", input)
		return true
	}
	fmt.Printf("%d is not Palindrome.\n", input)
	return false
}

func filter(arr []int, cond func(int) bool) []int {
	var result []int

	for _, v := range arr {
		if cond(v) {
			result = append(result, v)
		}
	}
	return result
}
