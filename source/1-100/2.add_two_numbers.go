package main

import "fmt"

/*
2. 两数相加

给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
*/

func main() {
	l1 := ListNode{
		2, &ListNode{
			4, &ListNode{
				3, nil,
			},
		},
	}
	l2 := ListNode{
		5, &ListNode{
			6, &ListNode{
				4, nil,
			},
		},
	}
	fmt.Println(addTwoNumbers(&l1, &l2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	return fmt.Sprintf("{Val: %v, Next: %v}", l.Val, l.Next)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := ListNode{0, nil}
	currentRes := &res

	for l1 != nil || l2 != nil {
		num := currentRes.Val

		if l1 != nil {
			num += l1.Val
		}

		if l2 != nil {
			num += l2.Val
		}

		if temp := num / 10; temp >= 1 {
			currentRes.Val = num % 10
			currentRes.Next = &ListNode{temp, nil}
		} else {
			currentRes.Val = num
		}

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

		if (l1 != nil || l2 != nil) && currentRes.Next == nil {
			currentRes.Next = &ListNode{0, nil}
		}
		currentRes = currentRes.Next
	}

	return &res
}
