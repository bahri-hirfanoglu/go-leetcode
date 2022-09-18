package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	node := head
	firstIndex := k
	lastIndex := (getCount(node) - k) + 1
	values := getValue(head, firstIndex, lastIndex)
	currentIndex := 1
	for node != nil {
		if firstIndex == currentIndex {
			node.Val = values["last"]
		}
		if lastIndex == currentIndex {
			node.Val = values["first"]
		}
		node = node.Next
		currentIndex += 1
	}
	return head
}

func getValue(head *ListNode, firstIndex, lastIndex int) map[string]int {
	currentIndex := 1
	arr := map[string]int{}
	for head != nil {
		if currentIndex == firstIndex {
			arr["first"] = head.Val
		}
		if currentIndex == lastIndex {
			arr["last"] = head.Val
		}
		if len(arr) == 2 {
			return arr
		}
		currentIndex += 1
		head = head.Next
	}
	return nil
}

func getCount(node *ListNode) int {
	temp := node
	count := 0
	for temp != nil {
		count++
		temp = temp.Next
	}
	return count
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
		Val: 100,
		Next: &ListNode{
			Val:  90,
			Next: nil,
		},
	}
	print(swapNodes(node, 2))
}
