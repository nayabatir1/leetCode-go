package main

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	map_1 := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
    temp := strings.Split(s,"")
	count := map_1[temp[len(temp) - 1]] 

	for i:=len(temp)-1; i>0; i--{
		if map_1[temp[i]] > map_1[temp[i-1]] {
			count -= map_1[temp[i-1]]
		}else {
			count += map_1[temp[i-1]]
		}
	}
	return count
}

func main()  {
	t := romanToInt("MCMLXXXVIII")
	fmt.Println(t)

}