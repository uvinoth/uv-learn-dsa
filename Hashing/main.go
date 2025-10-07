package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	set1 := make(map[int]bool)

	resultSet := make(map[int]bool)

	for _, num1 := range nums1 {
		set1[num1] = true
	}

	fmt.Println("set1", set1)

	for _, num2 := range nums2 {
		if set1[num2] {
			resultSet[num2] = true
		}
	}

	fmt.Println("resultSet", resultSet)

	result := make([]int, 0, len(resultSet))

	for val := range resultSet {
		result = append(result, val)
	}
	return result
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	output := intersection(nums1, nums2)
	fmt.Printf("Input 1: %v, Input 2: %v\nOutput: %v\n", nums1, nums2, output)
}
