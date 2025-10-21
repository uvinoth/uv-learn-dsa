package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func matrixArr() {
	matarr := [][]int{{1, 2}, {3, 4}, {5, 6}}

	for row := 0; row < len(matarr); row++ {
		for col := 0; col < len(matarr[row]); col++ {
			fmt.Println(matarr[row][col])
		}
		fmt.Println()
	}
}

func binarySearch() {
	arr := []int{-1, 0, 5}
	k := -1
	cnt := len(arr)
	low := 0
	high := cnt - 1

	for low <= high {
		mid := (low + high) / 2
		if k == arr[mid] {
			return
		} else if k < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("no record found")

}

// LeeCode 162
func findPeakElement(ar []int) int {
	cnt := len(ar)
	maxElemnt := ar[0]
	maxindex := 0

	for i := 1; i < cnt; i++ {
		if ar[i] > maxElemnt {
			maxElemnt = ar[i]
			maxindex = i
		}
	}
	return maxindex
}

func Rotateleft() {
	ar := []int{1, 2, 3, 4, 5}
	temp := ar[0]

	for i := 1; i < len(ar); i++ {
		ar[i-1] = ar[i]
	}
	ar[len(ar)-1] = temp

	fmt.Println(ar)
}

func reverseArr(arr []int, start, end int) {
	for start <= end {
		temp := arr[start]
		arr[start] = arr[end]
		arr[end] = temp
		start++
		end--
	}

}

func RotateArr(arr []int, k int) {
	k = k % len(arr)
	if k < 0 {
		k = k + len(arr)
	}

	reverseArr(arr, 0, len(arr)-1)
	reverseArr(arr, 0, k-1)
	reverseArr(arr, k, len(arr)-1)

	fmt.Println("arr", arr)
}

func ExtractInteger() {
	str := "1: Prakhar Agrawal, 2: Manish Kumar Rai, 3: Rishabh Gupta56"

	// re := regexp.MustCompile(`\d+`)
	// matches := re.FindAllString(str, -1)

	// var nums []int
	// for _, m := range matches {
	// 	n, _ := strconv.Atoi(m)
	// 	nums = append(nums, n)
	// }

	// fmt.Println(nums)

	// ***** Manual Way *******
	var nums []int
	current := ""

	for _, s := range str {

		if unicode.IsDigit(s) {
			current += string(s)
		} else if current != "" {
			n, _ := strconv.Atoi(current)
			nums = append(nums, n)
			current = ""
		}
	}

	if current != "" {
		n, _ := strconv.Atoi(current)
		nums = append(nums, n)
	}

	fmt.Println(nums)
}

func maxProfit() {
	ar := []int{7, 1, 5, 3, 6, 4}
	minSoFor := ar[0]
	profit := 0

	for i := 1; i < len(ar); i++ {
		cost := ar[i] - minSoFor
		profit = max(profit, cost)
		minSoFor = min(minSoFor, ar[i])
	}
	fmt.Println(profit)
}

func MajorityElement() {
	ar := []int{2, 2, 1, 1, 1, 2, 2}

	cnt := 0
	major := 0

	for _, n := range ar {
		if cnt == 0 {
			major = n
		} else if n == major {
			cnt++
		} else {
			cnt--
		}
	}

	fmt.Println(major)

}

func MajorityElementMap(ar []int) int {
	if len(ar) == 0 {
		return -1
	}

	major := make(map[int]int)

	for _, num := range ar {
		major[num]++
	}

	fmt.Println("major!!!!!", major)

	n := len(ar)
	for num, count := range major {
		if count > n/2 {
			return num
		}
	}

	return -1
}

// First, choose a candidate
// from the given set of elements if it is the same as the candidate element, increase the votes.
// Otherwise, decrease the votes if votes become 0, select another new element as the new candidate.
func majorityElementOptimal(nums []int) int {
	count := 0
	candidate := 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		} else if candidate == num {
			count++
		} else {
			count--
		}
	}

	finalCount := 0
	for _, num := range nums {
		if num == candidate {
			finalCount++
		}
	}

	if finalCount > len(nums)/2 {
		return candidate
	}

	return -1

}

func ProductofArray(ar []int) []int {
	n := len(ar)
	if n == 0 {
		return []int{}
	}

	// Initialize the result slice with the same size as the input array.
	res := make([]int, n)
	// ar := []int{1, 2, 3, 4}
	// First Pass: Calculate prefix products (left to right).
	// The prefix product starts at 1.
	prefixProduct := 1
	for i := 0; i < n; i++ {
		res[i] = prefixProduct
		// fmt.Println("reP", res)
		prefixProduct *= ar[i]
		// fmt.Println("prefixProduct", prefixProduct)
	}
	fmt.Println("resp", res)
	// Second Pass: Calculate suffix products (right to left) and combine with prefix products.
	// The suffix product starts at 1.
	suffixProduct := 1
	for i := n - 1; i >= 0; i-- {
		fmt.Println("resp1", res)
		res[i] *= suffixProduct
		fmt.Println("resS", res)
		suffixProduct *= ar[i]
	}

	return res
}

func main() {
	// arr := []int{22, 55, 77, 88, 99}
	// reverseArr(arr, 0, len(arr)-1)
	// matrixArr()
	// binarySearch()
	// arr := []int{22, 55, 77, 66, 99}
	// arrrSrt := []string{"vinoth", "Athiyan", "Agaran"}
	// sort.Ints(arr)
	// fmt.Println(arr)
	// sort.Strings(arrrSrt)
	// fmt.Println(arrrSrt)

	// ar := []int{1, 2, 1, 3, 5, 6, 4}

	// reult := findPeakElement(ar)
	// fmt.Println(reult)

	// Rotate array lift side
	// Rotateleft()
	// rarr := []int{1, 2, 3, 4, 5, 6, 7}
	// RotateArr(rarr, 3)

	// ExtractInteger()
	// maxProfit()
	// MajorityElement()

	// ar2 := []int{3, 3, 4, 2, 4, 4, 2, 4, 4}
	// result2 := MajorityElementMap(ar2)
	// fmt.Printf("The majority element is: %d\n", result2)

	// ar2 := []int{3, 3, 4, 2, 4, 4, 2, 4, 4}
	// result := majorityElementOptimal(ar2)

	// fmt.Printf("The majority element Optimal is: %d\n", result)

	// ProductofArray  ************
	ar := []int{1, 2, 3, 4}
	result := ProductofArray(ar)
	fmt.Println("ProductofArray", result)

}
