package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	head *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n1, _ := nodeToStr(l1)
	n2, _ := nodeToStr(l2)
	t := n1 + n2

	return arrToNode(strings.Split(reverse(strconv.Itoa(t)), ""))
}

func arrToNode(arr []string) *ListNode {
	list := &List{}
	for _, sum := range arr {
		num, _ := strconv.Atoi(sum)
		list.Insert(num)
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

func nodeToStr(l *ListNode) (int, error) {
	str := strconv.Itoa(l.Val)
	for l.Next != nil {
		str += strconv.Itoa(l.Next.Val)
		l = l.Next
	}
	return strconv.Atoi(str)
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	fmt.Println(addTwoNumbers(l1, l2))
}
