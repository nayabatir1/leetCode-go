package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
    n1,err1 := strconv.ParseInt(a, 2, 64)
    n2,err2 := strconv.ParseInt(b, 2, 64)

	fmt.Println(err1)

	if err1 == nil && err2 == nil {
		sum:= n1+n2

		return strconv.FormatInt(sum, 2)
	}

	return ""
}

func main(){
	t:= addBinary("10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
	"110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011")
	fmt.Println(t)
}