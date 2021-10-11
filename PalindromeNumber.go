package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPalindrome(x int) bool {
    str := strings.Split(strconv.Itoa(x),"")
	halfLength := len(str) / 2
	for  i:= 0; i<= halfLength;i++{
		if(str[i] != str[len(str) - 1 -i]){
			return false
		}
	}

	return true
}

func main(){
	x := 121
	fmt.Print(isPalindrome(x))
}