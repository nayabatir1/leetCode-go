package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	copy := strings.Split(s, " ")

	for i:=len(copy)-1 ; i>=0; i--{
		size:= len(copy[i])
		if size > 0 {
			return size
		}
	}

	return 0
}

func main() {
	t := lengthOfLastWord("   fly me   to   the moon  ")
	fmt.Println(t)
}