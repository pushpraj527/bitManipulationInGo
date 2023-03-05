/*
https://leetcode.com/problems/binary-search/
*/

package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	fmt.Println(search(arr, 9))
}

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		}
	}
	return -1
}
