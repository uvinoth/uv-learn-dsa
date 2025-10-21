package main

import "fmt"

// When given numbers from range 1, N

func cyclicsort(nums []int) {
	n := len(nums)
	i := 0

	for i < n {
		correct := nums[i] - 1
		if nums[i] != nums[correct] {
			nums[i], nums[correct] = nums[correct], nums[i]
		} else {
			i++
		}
	}

	for _, num := range nums {
		fmt.Print(" ", num)
	}
}

func missingNumber(nums []int) int {
	n := len(nums)
	exceptSum := n * (n + 1) / 2

	actualSum := 0
	for _, num := range nums {
		actualSum += num
	}

	return exceptSum - actualSum
}

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	i := 0

	for i < n {
		correct := nums[i] - 1
		if nums[i] != nums[correct] {
			nums[i], nums[correct] = nums[correct], nums[i]
		} else {
			i++
		}
	}
	fmt.Println("!!!!!", nums)
	result := []int{}
	for i, num := range nums {
		if num != i+1 {
			result = append(result, i+1)
		}
	}
	return result
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

	for i, num := range nums {
		if num != i+1 {
			return i + 1
		}
	}
	return n + 1
}

func findDuplicate(nums []int) int {
	i := 0
	n := len(nums)

	for i < n {
		correct := nums[i] - 1
		if nums[i] != nums[correct] {
			nums[i], nums[correct] = nums[correct], nums[i]
		} else {
			if i != correct {
				return nums[i]
			}
			i++
		}

	}

	return -1
}

func findAllDuplicates(nums []int) []int {
	i := 0
	n := len(nums)

	for i < n {
		correct := nums[i] - 1
		if nums[i] != nums[correct] {
			nums[i], nums[correct] = nums[correct], nums[i]
		} else {
			i++
		}
	}
	result := []int{}
	for i, num := range nums {
		if num != i+1 {
			result = append(result, num)
		}
	}
	return result
}

func findErrorNums(nums []int) []int {
	i := 0
	n := len(nums)

	for i < n {
		correct := nums[i] - 1
		if nums[i] != nums[correct] {
			nums[i], nums[correct] = nums[correct], nums[i]
		} else {
			i++
		}
	}

	for i, num := range nums {
		if num != i+1 {
			return []int{num, i + 1}
		}
	}
	return []int{}
}

func main() {
	// Problem 1: Missing Number  (LeetCode #268)
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))

	// Problem 2: Find All Numbers Disappeared in an Array (LeetCode #448)
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(nums)) // Output: [5 6]

	// Problem 3: First Missing Positive (LeetCode #41)
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))         // 3
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))     // 2
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12})) // 1

	// Problem 4: Find the Duplicate Number (LeetCode #287)
	arra := []int{3, 1, 3, 4, 2}
	fmt.Println("findDuplicate", findDuplicate(arra)) // Output: 3
	fmt.Println("findDuplicates")

	// Problem 5: Find All Duplicates in an Array LeetCode #442)
	numsArr := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findAllDuplicates(numsArr)) // Output: [2 3]

	// Problem 6: Set Mismatch LeetCode #645)
	numMis := []int{1, 2, 2, 4}
	fmt.Println(findErrorNums(numMis)) // Output: [2 3]

}
