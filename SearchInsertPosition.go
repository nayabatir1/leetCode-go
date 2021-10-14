package main

import "fmt"

func searchInsert(nums []int, target int) int {

	 halfSize := len(nums)/2
    
	if target > nums[halfSize]{
		for o:= halfSize; o<len(nums);o++{
			if target <= nums[o] {
				return o
			}
		}
	}else {
		for o:= 0; o<=halfSize;o++{
			if target <= nums[o] {
				return o
			}
		}
	}

	return len(nums)
}

func main() {
	t := searchInsert([]int{1,3,5}, 4)
	fmt.Println(t)
}