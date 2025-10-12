package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	set1 := make(map[int]bool)

	for _, num1 := range nums1 {
		set1[num1] = true
	}
	fmt.Println("set1", set1)

	resultSet := make(map[int]bool)
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

func singleNumberMap(nums []int) int {
	counts := make(map[int]int)

	for _, num := range nums {
		counts[num]++
	}

	for num, count := range counts {
		if count == 1 {
			return num
		}
	}

	fmt.Println("counts", counts)
	return 0
}

func distinctElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	set := make(map[int]bool)

	for _, num := range nums {
		set[num] = true
	}
	fmt.Println("set", set)
	return len(set)
}

func main() {
	// Intersection
	// nums1 := []int{4, 9, 5}
	// nums2 := []int{9, 4, 9, 8, 4}
	// output := intersection(nums1, nums2)
	// fmt.Printf("Input 1: %v, Input 2: %v\nOutput: %v\n", nums1, nums2, output)

	// Single Number Map
	nums := []int{4, 1, 2, 1, 2}
	output1 := singleNumberMap(nums)
	fmt.Printf("Input: %v, Output: %d\n", nums, output1)

	// Distinct element
	arr := []int{2, 2, 3, 2, 5}
	result := distinctElement(arr)
	fmt.Printf("Input: %v, Output: %d\n", arr, result)
}
