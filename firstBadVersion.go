/*
278. First Bad Version: https://leetcode.com/problems/first-bad-version/
*/
package main

import "fmt"

func main() {
	fmt.Println(firstBadVersion(5))
}

func isBadVersion(n int) bool {
	if n >= 2 {
		return true
	}
	return false
}

func firstBadVersion(n int) int {
	start := 1
	end := n
	for start < end {
		mid := (start + end) / 2
		if !isBadVersion(mid) {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}
