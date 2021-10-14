package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
    sum, ret := 0, math.MinInt16

	for _, v := range nums {
		if sum > 0 {
			sum += v
		}else {
			sum = v
		}
		if sum > ret {
			ret = sum
		}
	}
	return ret
}

func main(){
	t:= maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4})
	fmt.Println(t)
}