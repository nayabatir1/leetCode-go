package main
type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    m := &ListNode{}
    carry := 0
    n:=m
    for l1 != nil || l2!= nil {
        switch {
			case l1 == nil:{
				var sum ListNode 
				sum.Val= l2.Val+ carry
				if sum.Val >= 10 {
					carry = sum.Val / 10
					sum.Val %= 10
				}else {
					carry = 0
				}
				n.Next, l2 = &sum, l2.Next
			}
			case l2 == nil: {
				var sum ListNode 
				sum.Val= l1.Val+ carry
				if sum.Val >= 10 {
					carry = sum.Val / 10
					sum.Val %= 10
				}else {
					carry = 0
				}
				n.Next, l1 = &sum, l1.Next
			}
			default: {
				var sum ListNode 
				sum.Val = l1.Val + l2.Val + carry
				if sum.Val >= 10 {
					carry = sum.Val / 10
					sum.Val %= 10
				}else {
					carry = 0
				}
				n.Next = &sum
				l1 = l1.Next
				l2 = l2.Next
			}
    	}
    	n=n.Next
 	}
    if carry != 0 {
        var sum ListNode
        sum.Val = carry
        n.Next = &sum
    }
    
    return m.Next
}   