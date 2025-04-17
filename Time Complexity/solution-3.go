package main

import (
	"fmt"
	"slices"
)

func Solution(A []int) int {
	// sort the array
	slices.Sort(A)
	lo, hi := 0, len(A)-1
	for lo < hi {
		mid := int(lo + (hi-lo)/2)
		// If mid is odd, make it even
		if mid%2 == 1 {
			mid--
		}
		if A[mid] == A[mid+1] {
			lo = mid + 2
		} else {
			hi = mid
		}
	}
	return A[lo]
}

func main() {
	A := []int{1, 1, 2, 1000000001, 1000000001}
	B := []int{1, 3, 4, 5, 4, 3, 1}
	C := []int{1, 3, 4, 4, 3, 1, 2, 2, 6, 7, 7, 6, 100, 200, 269, 200, 100}
	D := []int{9, 3, 9, 3, 9, 7, 9}
	fmt.Println(Solution(A))
	fmt.Println(Solution(B))
	fmt.Println(Solution(C))
	fmt.Println(Solution(D))
}
