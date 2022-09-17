package main

import "fmt"

func searchInsert(nums []int, target int) int {
	index := -1
	for i, v := range nums {
		if v == target {
			return i
		}
		if target < v {
			index = i
			break
		}
	}
	if index == -1 {
		return len(nums)
	} else {
		return index
	}
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 2
	fmt.Println(searchInsert(nums, target))
}
