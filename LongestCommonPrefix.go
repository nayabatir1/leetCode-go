package main

import (
	"fmt"
	// "strings"
)

func longestCommonPrefix(strs []string) string {
    var prefix string

	smallest := len(strs[0])

	for i:=1;i<len(strs);i++ {
		if len(strs[i]) < smallest {
			smallest = len(strs[i])
		}
	}

	for i:= 0;i<smallest;i++{
		check := string([]rune(strs[0])[i])
		for j:=1 ; j<len(strs);j++{
			if check != string([]rune(strs[j])[i]) {
				check = ""
			}
		}
		if check != ""{
			prefix += check
		}else {
			return prefix
		}
	}

	return prefix
}

func main(){
	t := longestCommonPrefix([]string{"aa","ab"})
	fmt.Println(t)
}
