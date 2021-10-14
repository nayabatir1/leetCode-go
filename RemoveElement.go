package main

import "fmt"

func removeElement(nums []int, val int) int {
    offset := 0

	for _,v := range nums {
		if v != val {
			nums[offset] = v
			offset ++
		}
	}

	return len(nums[:offset])
}

func main(){
	t := removeElement([]int{1,2,3,4,4,5}, 4)
	fmt.Println(t)
}