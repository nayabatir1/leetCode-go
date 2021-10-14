package main

import (
	"fmt"
	"strings"
)

func strStr(haystack string, needle string) int {
	temp := strings.Split(haystack, needle)

	if len(temp) == 1 && needle != "" {
		return -1
	}
	if len(temp) > 1 {
		if len(temp[0]) != 0 {
			return len(temp[0])
		}
	}

    return 0
}

func main() {
	t := strStr("a", "")
	fmt.Println(t)
}