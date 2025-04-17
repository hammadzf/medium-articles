package main

import "fmt"

func Solution(A []int) int {
	unpaired := 0
	for _, element := range A {
		unpaired ^= element
	}
	return unpaired
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
