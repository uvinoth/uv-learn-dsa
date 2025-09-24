package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func reverseArr(arr []int, start, end int) {
	for start <= end {
		temp := arr[start]
		arr[start] = arr[end]
		arr[end] = temp
		start++
		end--
	}

}

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
	prifit := 0

	for i := 1; i < len(ar); i++ {
		cost := ar[i] - minSoFor
		prifit = max(prifit, cost)
		minSoFor = min(minSoFor, ar[i])
	}
	fmt.Println(prifit)
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
	maxProfit()

}
