package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(node *ListNode) *ListNode {
	head := node
	for head != nil && head.Next != nil {
		temp := head.Val
		head.Val = head.Next.Val
		head.Next.Val = temp
		head = head.Next.Next
	}
	return node
}

func print(node *ListNode) {
	_node := node
	for _node != nil {
		fmt.Println(_node.Val)
		_node = _node.Next
	}
}

func main() {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	print(swapPairs(node))
}
