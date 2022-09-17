package main

import (
	"fmt"
	"sort"
)

func removeElement(nums []int, val int) int {
	counter := 0
	for i, v := range nums {
		if v == val {
			counter += 1
			nums[i] = '_'
		}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return len(nums) - counter
}
func main() {
	nums := []int{3, 2, 2, 3}
	target := 3
	fmt.Println(removeElement(nums, target))
}
