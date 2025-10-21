package main

import (
	"fmt"
	"math"
	"sort"
)

func bubbleSort() {
	fmt.Println("Bubblen  Sort")
	arr := []int{15, 16, 6, 8, 5}

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println(arr)
}

func selectionSort() {
	fmt.Println("Selection  Sort")
	arr := []int{15, 16, 6, 8, 5}
	n := len(arr)

	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}

		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}

	fmt.Println("Selection", arr)
}

func insertionSort() {
	fmt.Println("Insertion  Sort")
	ar := []int{5, 4, 10, 1, 6, 2}
	n := len(ar)

	for i := 1; i < n; i++ {
		temp := ar[i]
		j := i - 1

		for j >= 0 && ar[j] > temp {
			ar[j+1] = ar[j]
			j--
		}

		ar[j+1] = temp

	}

	fmt.Println("Insertion", ar)

}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	fmt.Println("mid", mid)

	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	// fmt.Println("left", arr[:mid])
	// fmt.Println("right", arr[mid:])

	return mearge(left, right)
}

func mearge(left, right []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// The append function adds all those remaining elements to the result slice.
	// The ... (ellipsis) expands the slice into individual elements.\\
	// Without it, Go would try to append the entire slice as a single element (which isn’t allowed).
	// With ..., Go appends each element of left[i:] individually.

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func quick_sort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high)
		fmt.Print(p)
		quick_sort(arr, low, p-1)  // left side
		quick_sort(arr, p+1, high) // right side
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high] // choose last element as pivot
	i := low - 1
	fmt.Println("iii", i)

	for j := low; j < high; j++ {
		if arr[j] < pivot { // strictly less than pivot
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)/2]
	left := []int{}
	right := []int{}

	for i, val := range arr {
		if i == len(arr)/2 {
			continue // skip pivot itself
		}

		if val <= pivot {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}

	sortLeft := quickSort(left)
	sortRight := quickSort(right)

	return append(append(sortLeft, pivot), sortRight...)
}

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	result := [][]int{}
	maxDiff := math.MaxInt

	// Step 2: Find the minimum absolute difference
	for i := 1; i < len(arr); i++ {
		diff := arr[i] - arr[i-1]
		if diff < maxDiff {
			maxDiff = diff
		}
	}

	// Step 3: Collect all pairs having the minimum difference
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] == maxDiff {
			result = append(result, []int{arr[i-1], arr[i]})
		}
	}
	return result
}

func wiggleSort(nums []int) {
	n := len(nums)
	result := append([]int{}, nums...)
	sort.Ints(result)

	mid := (n + 1) / 2

	left := result[:mid]
	right := result[mid:]

	fmt.Println("left", left)
	fmt.Println("right", right)

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Println("!!!", left[len(left)-1])
			nums[i] = left[len(left)-1]
			fmt.Println("left if", left[:len(left)-1])
			left = left[:len(left)-1]

		} else {
			fmt.Println("!!!", right[len(right)-1])
			nums[i] = right[len(right)-1]
			fmt.Println("right if", right[:len(right)-1])
			right = right[:len(right)-1]
		}
	}
}

func sortColors(nums []int) []int {
	low, mid, high := 0, 0, len(nums)-1

	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
	return nums
}

func cyclicsort(nums []int) {
	n := len(nums)
	i := 0

	for i < n {
		corrct := nums[i] - 1
		if nums[i] != nums[corrct] {
			nums[i], nums[corrct] = nums[corrct], nums[i]
		} else {
			i++
		}
	}

	for _, num := range nums {
		fmt.Print(" ", num)
	}
}

func firstMissingPositive(nums []int) int {
	n := len(nums)
	i := 0

	for i < n {
		correct := nums[i] - 1
		if nums[i] > 0 && nums[i] <= n && nums[i] != nums[correct] {
			nums[i], nums[correct] = nums[correct], nums[i]
		} else {
			i++
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}

func main() {
	// Bubble sort
	// bubbleSort()

	// Selection sort
	// Find the smallest element and swap it with the first element
	// selectionSort()

	// Insertion sort
	// Pick an element and place it in the correct position in the sorted part
	// insertionSort()

	// Merge Sort
	// arr := []int{38, 27, 43, 3, 9, 82, 10}
	// fmt.Println("Original:", arr)
	// sorted := mergeSort(arr)
	// fmt.Println("Sorted:", sorted)

	// Quick sort
	// - How it works:
	// - Choose a value in the array to be the pivot element.
	// - Order the rest of the array so that lower values than the pivot element are on the left, and higher values are on the right.
	// - Swap the pivot element with the first element of the higher values so that the pivot element lands in between the lower and higher values.
	// - Do the same operations (recursively) for the sub-arrays on the left and right side of the pivot element.

	// Quick Sort is another divide-and-conquer algorithm:
	// 	Pick a pivot element.
	// 	Partition the array into:
	// 	- elements less than pivot,
	// 	- elements equal to pivot, and
	// 	- elements greater than pivot.
	// 	Recursively sort the left and right partitions.
	//  Combine the results.

	// fmt.Println("Quick sort")
	// ar := []int{7, 6, 10, 5, 9, 2, 1, 15, 7}
	// fmt.Println("Before sorting:", ar)
	// quickSort(ar, 0, len(ar)-1)
	// fmt.Println("After sorting:", ar)

	// nums := []int{10, 7, 8, 9, 1, 5}
	// fmt.Println("Original:", nums)
	// sorted := quickSort(nums)
	// fmt.Println("Sorted:", sorted)

	// Minimum Absolute Difference - 1200
	// fmt.Println(minimumAbsDifference([]int{4, 2, 1, 3}))                     // [[1 2] [2 3] [3 4]]
	// fmt.Println(minimumAbsDifference([]int{1, 3, 6, 10, 15}))                // [[1 3]]
	// fmt.Println(minimumAbsDifference([]int{3, 8, -10, 23, 19, -4, -14, 27})) // [[-14 -10] [19 23] [23 27]]

	// Problems  324 : Wiggle Sort II
	// Sort the array.
	// Split into two halves:
	// Smaller half → goes to even indices.
	// Larger half → goes to odd indices.
	// Fill from the end of both halves to avoid equal elements causing issues.

	// nums1 := []int{1, 5, 1, 1, 6, 4}
	// wiggleSort(nums1)
	// fmt.Println(nums1) // e.g. [1 6 1 5 1 4]

	// nums2 := []int{1, 3, 2, 2, 3, 1}
	// wiggleSort(nums2)
	// fmt.Println(nums2) // e.g. [2 3 1 3 1 2]

	// Problems  75 :Sort Colors
	// We maintain three pointers:
	// low: marks the boundary where 0s end.
	// mid: current element under consideration.
	// high: marks the boundary where 2s begin.

	// We iterate while mid <= high:
	// If nums[mid] == 0: swap with low, increment both low and mid.
	// If nums[mid] == 1: just move mid forward.
	// If nums[mid] == 2: swap with high, decrement high (don’t increment mid yet).

	// nums := []int{2, 0, 2, 1, 1, 0}
	// sortColors(nums)
	// fmt.Println(nums) // Output: [0 0 1 1 2 2]

	// Cyclic sort
	// arr := []int{4, 3, 2, 1, 5}
	// cyclicsort(arr)

	// Problem: First Missing Positive (LeetCode #41)
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))         // 3
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))     // 2
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12})) // 1
}
