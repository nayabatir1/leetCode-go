package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
    
    n := head
    
    for l1 != nil || l2 != nil {
        switch {
            case l1 == nil : n.Next, l2 = l2, l2.Next
            case l2 == nil : n.Next , l1 = l1 , l1.Next
            case l1.Val > l2.Val : n.Next , l2 = l2 , l2.Next
            default: n.Next, l1 = l1, l1.Next
        }
        n = n.Next
    }
    return head.Next
}

func main(){
		
}