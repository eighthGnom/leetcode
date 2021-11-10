package third_day

import (
	"fmt"
)

/*Given an unsorted integer array nums, return the smallest missing positive integer.

You must implement an algorithm that runs in O(n) time and uses constant extra space.*/

func firstMissingPositive1(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[0] < 1 {
			nums[0] = nums[len(nums)-1]
			nums[len(nums)-1] = 0
			nums = nums[:len(nums) - 1]
			continue
		}else  if nums[len(nums)-1] < 1 {
			nums[len(nums)-1] = nums[0]
			nums[0] = 0
			nums = nums[1:]
		}
		x := nums[i]
		j := i
		for ; j > 0 && nums[j-1] > x; j-- {
			nums[j] = nums[j-1]

		}
		nums[j] = x
	}
	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1

}



func test_1() {

	fmt.Println(firstMissingPositive1([]int{1,2,0}))
}


func firstMissingPositive2(nums []int) int {
	m := map[int]bool{}
	max := 0
	for _, i := range nums {
		if i > max {
			max = i
		}
		if i > 0 {
			m[i] = true
		}
	}
	for i := 1; i <= max; i++{
		if _, j := m[i]; !j {
			return i
		}
	}
	return max + 1
}

func test_2() {
	fmt.Println(firstMissingPositive2([]int{1,2,0}))
}

