package main

import (
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	head *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	list1_arr := nodeToArr(list1)
	list2_arr := nodeToArr(list2)
	list := append(list1_arr, list2_arr...)
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
	last_list := arrToNode(list)
	return last_list
}

func nodeToArr(list *ListNode) []int {
	arr := []int{}
	if list == nil {
		return arr
	}
	arr = append(arr, list.Val)
	for list.Next != nil {
		arr = append(arr, list.Next.Val)
		list = list.Next
	}
	return arr
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

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	mergeTwoLists(list1, list2)

}
