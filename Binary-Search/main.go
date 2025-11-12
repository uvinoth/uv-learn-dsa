package main

import (
	"fmt"
	"sort"
)

func binarySearch(nums []int, target int) int {
	// If check array is sorted or not
	n := len(nums)
	start, end := 0, n-1

	for start <= end {
		mid := (start + end) / 2

		if target == nums[mid] {
			return mid
		} else if nums[mid] < target { //If target grater the mid ignore the left side
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return -1
}

func search(nums []int, target int) int {
	start, end := 0, len(nums)-1

	for start <= end {
		mid := start + (end-start)/2 //Out of interger using BODMAS algo

		if nums[mid] == target {
			return mid
		}
		// Determined sorted array
		if nums[start] <= nums[mid] {
			if nums[start] <= target && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}

	return -1
}

// Search Element in Rotated Sorted Array II
func searchII(nums []int, target int) bool {
	n := len(nums)
	left, right := 0, n-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return true
		}

		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
			continue
		}

		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])

	low := 0
	high := m*n - 1

	for low <= high {
		mid := low + (high-low)/2
		row := mid / 2
		col := mid % 2
		minValue := matrix[row][col]

		if target == minValue {
			return true
		} else if target < minValue {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

func lowerBound(arr []int, target int) int {
	n := len(arr)
	left := 0
	right := n - 1
	res := n

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] >= target {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}

func upperBound(arr []int, target int) int {
	n := len(arr)
	left := 0
	right := n - 1
	res := n
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] > target {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}
func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right, res := 0, n-1, n

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] >= target {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}

func searchRange(nums []int, target int) []int {
	return []int{findFirst(nums, target), findLast(nums, target)}
}

func findFirst(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	firstEle := -1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			firstEle = mid
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return firstEle
}

func findLast(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	last := -1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			last = mid
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return last
}

func findMin(nums []int) int {
	n := len(nums)
	left, right := 0, n-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}

func findArrayRotated(nums []int) int {
	n := len(nums)
	left, right := 0, n-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

// If you ever get confused:
// If mid is even, compare with mid+1
// If mid is odd, compare with mid-1
// The first mismatch marks where the single element lies
func singleNonDuplicate(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if nums[0] != nums[1] {
		return nums[0]
	}
	if nums[n-1] != nums[n-2] {
		return nums[n-1]
	}

	left, right := 1, n-2

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] != nums[mid-1] && nums[mid] != nums[mid+1] {
			return nums[mid]
		}

		if (mid%2 == 1 && nums[mid] == nums[mid-1]) || (mid%2 == 0 && nums[mid] == nums[mid+1]) {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return -1
}

func findPeak(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func mySqrt(x int) int {
	left, right := 0, x

	for left <= right {
		mid := left + (right-left)/2
		val := mid * mid

		if val <= x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}

func Sqrt(mid, n int) int {
	ans := 1
	for n > 0 {
		if n%2 == 1 {
			ans = ans * mid
			n = n - 1
		} else {
			mid = mid * mid
			n = n / 2
		}
	}

	return ans
}

func NthSqrt(m int) int {
	left, right := 1, m
	for left <= right {
		mid := left + (right-left)/2
		midN := Sqrt(mid, 3)
		if midN == m {
			return mid
		} else if midN < m {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func aggressiveCows(stalls []int, cows int) int {
	sort.Ints(stalls)
	left, right := 1, stalls[len(stalls)-1]-stalls[0]
	res := -1

	for left <= right {
		mid := left + (right-left)/2

		if canPlace(stalls, cows, mid) {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return res
}

func canPlace(stalls []int, cows int, mid int) bool {
	lastPos := stalls[0]
	count := 1

	for i := 1; i < len(stalls); i++ {
		if stalls[i]-lastPos >= mid {
			count++
			lastPos = stalls[i]
		}
	}
	if count == cows {
		return true
	}

	return false
}

func allocateBooks(books []int, std int) int {
	n := len(books)
	if std > n {
		return -1
	}

	low, high := 0, 0

	for _, book := range books {
		if book > low {
			low = book
		}
		high += book
	}

	result := high
	for low <= high {
		mid := low + (high-low)/2
		// fmt.Println("mid", mid)
		if canAllocation(books, std, mid) {
			result = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return result
}

func canAllocation(books []int, stud, maxpage int) bool {
	allocateStud := 1
	allocatePage := 0

	for _, page := range books {
		if page > maxpage {
			return false
		}

		if page+allocatePage > maxpage {
			allocateStud++
			allocatePage = page
			if allocateStud > stud {
				return false
			}
		} else {
			allocatePage += page
		}
	}

	return true
}

func main() {
	// problem : binarySearch (leecode 704)
	// nums := []int{-1, 0, 3, 5, 9, 12}
	// target := 9
	// fmt.Println(binarySearch(nums, target))

	// problem : Search in a rotated Sorted Array (leecode 33)
	// nums := []int{4, 5, 6, 7, 0, 1, 2}
	// target := 0
	// fmt.Println(search(nums, target)) // Output: 4

	// matrix := [][]int{
	// 	{1, 3, 5, 7},
	// 	{10, 11, 16, 20},
	// 	{23, 30, 34, 60},
	// }
	// target := 3
	// fmt.Println(searchMatrix(matrix, target)) // Output: true

	// problem BS2: Implement Lower Bound and Upper Bound (Not leeCode)

	// arr := []int{1, 2, 4, 4, 5, 7}
	// x := 4
	// // Definition:
	// // The smallest index i such that arr[i] >= x.
	// fmt.Println(lowerBound(arr, x))
	// // Definition:
	// // The smallest index i such that arr[i] > x.
	// fmt.Println(upperBound(arr, x))

	// problem BS2: Search Insert Position (#leeCode 35)
	// fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	// fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	// fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))

	// // problem BS3: First and Last Occurrences in Array
	// nums := []int{5, 7, 7, 8, 8, 10}
	// target := 8
	// fmt.Println(searchRange(nums, target)) // Output: [3 4]

	// problem BS6: Minimum in Rotated Sorted Array (#leeCode 35)
	// nums := []int{4, 5, 6, 7, 0, 1, 2}
	// fmt.Println(findMin(nums)) //  Output: 0
	// // Find out how many times array has been rotated
	// fmt.Println(findArrayRotated(nums))

	// problem BS8 : Single Element in a Sorted Array (#leeCode 540)
	// nums := []int{1, 1, 2, 3, 3, 4, 4, 8, 8}
	// fmt.Println(singleNonDuplicate(nums)) // Output: 2

	// problem BS8 : Find Peak Element (#leeCode 162)
	// nums := []int{1, 2, 1, 3, 5, 6, 4}
	// fmt.Println(findPeak(nums)) // Output: 5

	// // problem BS10 : Finding Sqrt of a number using Binary Search (#leeCode 69)
	// x := 28
	// fmt.Println(mySqrt(x))

	// problem BS17 Aggressive cows - (Min dist between cows is max)
	// stalls := []int{1, 2, 8, 4, 9}
	// cows := 3
	// fmt.Println(aggressiveCows(stalls, cows)) // Output: 3

	// Problem Allocate Books
	//  - Max no of pages allocated is minimum.
	// 		- Each student get at least one book.
	// 		- Each book should be a allocated to only one student.
	// 		- Book allocation should be in a contiguous manner.

	books := []int{12, 34, 67, 90}
	students := 2
	fmt.Println(allocateBooks(books, students)) // Output: 113

}
