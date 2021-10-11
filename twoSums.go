// https://leetcode.com/problems/two-sum/
package main

import "fmt"

func twoSum(nums []int, target int) []int {
    map_1 := make(map[int]int)

	for i:=0;i<len(nums);i++{
		temp := target - nums[i]
		for j:=0; j<i;j++{
			if(map_1[nums[j]] == 0){
				map_1[nums[j]] = j
			}
		}
		fmt.Println(map_1, temp, map_1[temp])

		if _, ok := map_1[temp]; ok {
			return []int{i, map_1[temp]}
		}
	}
	return []int{}
}

func main(){
	t := twoSum([]int{2,7,11,15}, 9)
	fmt.Print(t)
}