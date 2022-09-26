package main

import (
	"fmt"
	"sort"
)

type List struct {
	head *ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var result []int
	for _, node := range lists {
		temp := node
		for temp != nil {
			result = append(result, temp.Val)
			temp = temp.Next
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return arrToNode(result)
}

func arrToNode(arr []int) *ListNode {
	list := &List{}
	for _, sum := range arr {
		list.Insert(sum)
	}
	return list.head
}

func (l *List) Insert(val int) {
	list := &ListNode{Val: val, Next: nil}
	if l.head == nil {
		l.head = list
	} else {
		p := l.head
		for p.Next != nil {
			p = p.Next
		}
		p.Next = list
	}
}

func print(node *ListNode) {
	_node := node
	for _node != nil {
		fmt.Println(_node.Val)
		_node = _node.Next
	}
}
func main() {
	var list []*ListNode
	item1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	item2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	item3 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  6,
			Next: nil,
		},
	}
	list = append(list, item1)
	list = append(list, item2)
	list = append(list, item3)
	mergeKLists(list)
}