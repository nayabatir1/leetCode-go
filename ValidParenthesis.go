package main

import (
	"fmt"
)

func isValid(s string) bool {
    gate := true
	m_a_p := map[string]string{")": "(", "}": "{", "]": "["}

	stack := []string{}
	for _, char := range s {
		if el,ok := m_a_p[string(char)]; ok{
			if len(stack) == 0{
				return false
			}else if stack[len(stack) -1 ] == el {
				stack = stack[: len(stack) - 1]
			}else {
				return false
			}
		}else {
			stack = append(stack, string(char))
		}
	}
	if len(stack) != 0 {
		gate = false
	}

	return gate
}

func main(){
	t := isValid("(])")
	fmt.Println(t)
}