package main

import "fmt"

func removeDuplicates(nums []int) int {
    temp := make(map[int]int)

	offset:=0

	for _, v := range nums {
		if _, ok := temp[v]; !ok {
			temp[v] = offset
			nums[offset] = v
			offset++
		}
	}

	return len(nums[:offset])
}

func main() {
	t := removeDuplicates([]int{1,1,2,3,4})
	fmt.Println(t)
}