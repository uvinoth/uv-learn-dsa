package main

import "fmt"

func SubArray() {
	k := 3
	ar := []int{1, 2, 3, 4, 1}
	cnt := 0
	n := len(ar)

	for i := 0; i < n; i++ {
		fmt.Println("outer", ar[i])
		sum := 0
		for j := i; j < n; j++ {
			sum += ar[j]
			fmt.Println("sum", sum)
			if sum == k {
				cnt++
			}
		}

	}

	fmt.Println(cnt)
}

func SubarraySumEqualsK(arr []int, k int) int {
	// map[sum]count - stores the frequency of each prefix sum encountered
	prefixSumCounts := make(map[int]int)

	// Base case: A prefix sum of 0 appears once (before index 0)
	prefixSumCounts[0] = 1

	currentSum := 0
	count := 0

	for _, num := range arr {
		currentSum += num

		// The sum we are looking for: previousSum = currentSum - k
		// If this previous sum exists, it means a subarray ending at the current
		// index sums to k.
		targetPrevSum := currentSum - k

		// Add the count of how many times this targetPrevSum has appeared
		count += prefixSumCounts[targetPrevSum]

		// Record the current prefix sum for future calculations
		prefixSumCounts[currentSum]++
	}

	return count
}

// func MaxiSubarray(arr []int) int {
// 	maxSum := 0
// 	count := len(arr)
// 	for i := 0; i < count; i++ {
// 		sum := 0
// 		for j := i; j < count; j++ {
// 			sum += arr[j]
// 			maxSum = max(sum, maxSum)
// 			// if sum > maxSum {
// 			// 	maxSum = sum
// 			// }
// 		}
// 	}
// 	return maxSum
// }

// Standard Kadane's Algorithm
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	sum := 0

	for _, num := range nums {
		if sum >= 0 {
			sum += num
		} else {
			sum = num
		}
		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}

func maxSubArraySlidingWindow(arr []int, k int) int {
	if len(arr) < k {
		return 0
	}
	currentSum := 0
	for i := 0; i < k; i++ {
		currentSum += arr[i]
	}
	fmt.Println("currentSum", currentSum)
	maxSum := currentSum
	fmt.Println("maxSum", maxSum)
	for j := k; j < len(arr); j++ {
		currentSum = currentSum - arr[j-k] + arr[j]

		maxSum = max(currentSum, maxSum)
	}
	return maxSum
}

func findMaxLength(nums []int) int {
	maxlen := 0
	pSum := 0
	sumIndexMap := make(map[int]int)
	sumIndexMap[0] = -1

	for i, num := range nums {
		if num == 1 {
			pSum++
		} else {
			pSum--
		}

		if prevIndex, found := sumIndexMap[pSum]; found {
			curVal := i - prevIndex
			if curVal > maxlen {
				maxlen = curVal
			}
		} else {
			sumIndexMap[pSum] = i
		}
	}
	return maxlen
}

func main() {
	// ******** SubarraySumEqualsK Start ********
	// k := 3
	// ar := []int{1, 2, 3, 4, 1}
	// result := SubarraySumEqualsK(ar, k)
	// fmt.Println(result)
	// ******** SubarraySumEqualsK End ********

	// ******** Maximum subarray Start ********
	// Standard Kadane's Algorithm
	// nums2 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	// fmt.Println("Result for example:", maxSubArray(nums2))
	// ******** Maximum subarray end ********

	// ******** Maximum sum subarray of size k Start ********
	// ar := []int{2, 9, 31, -4, 21, 7}
	// k := 3
	// result := maxSubArraySlidingWindow(ar, k)
	// fmt.Println("Maximum contiguous subarray sum of size k (O(n) solution):", result)
	// ******** Maximum sum subarray of size k end ********

	// Contiguous Array
	arr := []int{0, 1, 1, 1, 1, 1, 0, 0, 0}
	result := findMaxLength(arr)
	fmt.Println("Contiguous Array", result)

}
