package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		subtarget := target - num
		if j, found := m[subtarget]; found {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}

func twoPointerApproch(nums []int, target int) string {
	left, right := 0, len(nums)-1
	sort.Ints(nums)
	for left <= right {
		sum := nums[left] + nums[right]
		if target == sum {
			return "Yes"
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return "No"
}

// Condition i+j+k = 0
// Sort the given array
// Fixed i value constant and skip the duplicate
// Difine the two pointer j,k
// Sum the all the 3 value i,j,k
// Check condition all 3 value equal to 0
// If equal to sum=0, append the 3 value in response
// Skip the both duplicate value j and k
// If sum < 0 increase the j value
// if sum > 0 decrease the k value
func threeSum(nums []int) [][]int {
	// Sort the nums
	sort.Ints(nums)
	// Define the result
	result := [][]int{}

	for i := 0; i < len(nums)-1; i++ {
		// Check the duplicate i value
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// define left anf right
		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// skip the duplicate left and right
				for left < right && nums[left] == nums[left+1] {
					left++
				}

				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	n := len(nums)

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left := j + 1
			right := n - 1

			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return result
}

// Observation
// We need minimum 3 or more blocks
// Descending blocks can not store water
// Ascending blocks can not store water
// How to calculate trap water?
// - Find lb and rb
// - Find water level (minimum of lb,rb). find the leftMax and rightMax
// - Find traped water (water level - height of bar)
// - Add all traped water and return the result
func trap(height []int) int {
	left, right := 0, len(height)-1
	leftmax, rightMax := 0, 0
	total := 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftmax {
				leftmax = height[left]
			} else {
				total += leftmax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				total += rightMax - height[right]
			}
			right--
		}
	}

	return total
}

// Observation
//
//	Find the amount of water a container can store
//
// Using two pointer
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left < right {
		currentHeight := min(height[left], height[right])
		width := right - left
		area := currentHeight * width
		maxArea = max(maxArea, area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

// Observation
func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	left, right := 0, n-1
	pos := n - 1

	for left <= right {
		leftSq := nums[left] * nums[left]
		rightSq := nums[right] * nums[right]

		if leftSq > rightSq {
			result[pos] = leftSq
			left++
		} else {
			result[pos] = rightSq
			right--
		}
		pos--
	}

	return result
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandPalindrome(s, i, i)
		len2 := expandPalindrome(s, i, i+1)
		pLen := max(len1, len2)

		if pLen > end-start {
			start = i - (pLen-1)/2
			end = i + pLen/2
		}
	}

	return s[start : end+1]
}

func expandPalindrome(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func main() {
	// Problem:Two Sum (Leecode 1)
	// nums := []int{2, 7, 11, 15}
	// target := 9
	// result := twoSum(nums, target)
	// fmt.Println("Indices:", result) // Output: [0, 1]
	// val := twoPointerApproch(nums, target)
	// fmt.Println("val:", val)

	// Problem: 3Sum (Leecode 15) Time Complexity = O(n²), Space Complexity: O(1)
	// nums := []int{-1, 0, 1, 2, -1, -4}
	// fmt.Println(threeSum(nums))

	// Problem: 4Sum (Leecode 15) Time Complexity = O(n3), Space Complexity: O(1)
	// nums := []int{1, 0, -1, 0, -2, 2}
	// target := 0
	// fmt.Println(fourSum(nums, target))

	//  Problem: Trapping Rain Water (Leecode 2)
	// height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	// fmt.Println(trap(height))

	// Problem: Container With Most Water (Leecode 11)
	// height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	// fmt.Println(maxArea(height))

	// Problem: Squares of a Sorted Array (Leecode 977)
	// fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10})) // [0 1 9 16 100]
	// fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11})) // [4 9 9 49 121]

	// Problem: Longest Palindromic Substring (Leecode 5)
	fmt.Println(longestPalindrome("babad")) // "bab" or "aba"
	// fmt.Println(longestPalindrome("cbbd"))  // "bb"
	// fmt.Println(longestPalindrome("a"))     // "a"
	// fmt.Println(longestPalindrome("ac"))    // "a"
}
