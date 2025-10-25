package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func isSubsequence(s, t string) bool {
	i, j := 0, 0
	rs := []rune(s)
	rt := []rune(t)

	for i < len(rs) && j < len(rt) {
		if rs[i] == rt[j] {
			i++
		}
		j++
	}
	fmt.Println(i)
	return i == len(rs)
}

func reverseString(s string) string {
	runes := []rune(s)
	fmt.Println("runes", runes)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func sortString(str string) string {
	r := []rune(str)

	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	return string(r)

}

func strCount() int {
	str := "This is Athiyan and Agaran"
	target := 'a'
	count := 0

	s := []rune(str)

	for i := 0; i < len(s); i++ {
		if s[i] == target {
			count++
		}
	}
	fmt.Println("count", count)
	return count
}

func removeChara() {
	str := "This is Athiyan and Agaran"
	target := 'a'

	// newStr := strings.Replace(str, target, ".", 2)

	// fmt.Println(newStr)
	var srNew []rune
	// for i := 0; i < len(sr); i++ {
	for _, ch := range str {
		if ch != target {
			srNew = append(srNew, ch)
		}
	}

	resultString := string(srNew)
	fmt.Println("srNew", resultString)
}

func capitalizeTitle(title string) string {
	words := strings.Split(title, " ")
	fmt.Println("words", words)

	for i, word := range words {
		n := len(word)

		lowerWord := strings.ToLower(word)

		if n <= 2 {
			words[i] = lowerWord
		} else {
			firstWord := string(lowerWord[0])
			restWord := lowerWord[1:]
			words[i] = strings.ToUpper(firstWord) + restWord
		}
	}

	return strings.Join(words, " ")
}

func isPalindrome(s string) bool {
	var filtered []rune

	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			filtered = append(filtered, unicode.ToLower(ch))
		}
	}

	i, j := 0, len(filtered)-1

	for i < j {
		if filtered[i] != filtered[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func reverseWords(s string) string {
	words := strings.Fields(s)
	fmt.Println("words", words)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

func compress(chars []byte) int {
	write := 0 // Position to write compressed data
	read := 0  // Position to read original data

	for read < len(chars) {
		char := chars[read]
		count := 0
		fmt.Println("char1", string(char))
		fmt.Println("chars!!!!", string(chars[read]))
		// Count occurrences of the same character
		for read < len(chars) && chars[read] == char {
			read++
			count++
		}

		fmt.Println("char2", string(char))
		// Write the character
		fmt.Println("write1", write)
		chars[write] = char
		write++

		fmt.Println("chars!!!! #######", string(chars))

		fmt.Println("count", count)

		// Write the count (if > 1)
		if count > 1 {
			for _, c := range strconv.Itoa(count) {
				fmt.Println("ccccccc", string(c))
				chars[write] = byte(c)
				write++
			}
		}
	}

	return write
}

func reverseOnlyLetters(s string) string {
	str := []rune(s)
	left, right := 0, len(str)-1

	for left < right {
		if !unicode.IsLetter(str[left]) {
			left++
			continue
		}

		if !unicode.IsLetter(str[right]) {
			right--
			continue
		}

		str[left], str[right] = str[right], str[left]

		left++
		right--
	}
	return string(str)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for _, s := range strs[1:] {

		for len(prefix) > 0 && len(s) >= len(prefix) && s[:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1]
		}

		for len(prefix) > len(s) {
			prefix = prefix[:len(prefix)-1]
		}

		if prefix == "" {
			return ""
		}
	}

	return prefix
}

func longestPalindrome(s string) int {
	counts := make(map[rune]int)

	for _, ch := range s {
		counts[ch]++
	}

	length := 0
	oddFound := false

	for _, c := range counts {
		if c%2 == 0 {
			length += c
		} else {
			length += c - 1
			oddFound = true
		}
	}

	if oddFound {
		length++
	}
	return length
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[rune]int)
	for _, ch := range s {
		count[ch]++
	}

	for _, ch := range t {
		count[ch]--
		if count[ch] < 0 {
			return false
		}
	}

	return true
}

func main() {
	// Problem 1
	// isSubsequence
	// fmt.Println(isSubsequence("ace", "abcde")) // true
	// fmt.Println(isSubsequence("aec", "abcde")) // false

	// Problem 2
	// reverseString
	// fmt.Println(reverseString("hello"))
	// fmt.Println(reverseString("go❤️"))

	// Problem 3
	// sortString
	// input := "golang"
	// sorted := sortString(input)
	// fmt.Printf("Original: %s\nSorted: %s\n", input, sorted)

	// Problem 4
	// strCount()

	// Problem 5
	// removeChara()

	// Problem 6
	// CapitalizeTitle
	// title1 := "capiTalIze tHe titLe"
	// fmt.Printf("Input: \"%s\"\nOutput: \"%s\"\n\n", title1, capitalizeTitle(title1))

	// Problem 7
	// isPalindrome
	// s := "A man, a plan, a canal: Panama"
	// result := isPalindrome(s)
	// fmt.Println(result)

	// Problem 8
	// s := "  hello world  "
	// result := reverseWords(s)
	// fmt.Println(result)

	// Problem 9
	// Compress
	// ex1 := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	// n1 := compress(ex1)
	// fmt.Println(n1, string(ex1[:n1])) // Output: 6, "a2b2c3"

	// Problem 10
	// ex2 := []byte{'a'}
	// n2 := compress(ex2)
	// fmt.Println(n2, string(ex2[:n2])) // Output: 1, "a"

	// Problem 11
	// ex3 := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	// n3 := compress(ex3)
	// fmt.Println(n3, string(ex3[:n3])) // Output: 4, "ab12"

	// Problem 12
	// reverse Only Letters
	// fmt.Println(reverseOnlyLetters("ab-cd"))                // Output: "dc-ba"
	// fmt.Println(reverseOnlyLetters("a-bC-dEf-ghIj"))        // Output: "j-Ih-gfE-dCba"
	// fmt.Println(reverseOnlyLetters("Test1ng-Leet=code-Q!")) // Output: "Qedo1ct-eeLg=ntse-T!"

	// Problem 13
	// longest Common Prefix
	// fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) // Output: "fl"
	// fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))

	// Problem 14
	// longest Palindrome
	fmt.Println(longestPalindrome("abccccdd")) // Output: 7
	fmt.Println(longestPalindrome("a"))        // Output: 1
	fmt.Println(longestPalindrome("Aa"))       // Output: 1

	// Problem 14  Anagram
	fmt.Println(isAnagram("anagram", "nagaram")) // true
	fmt.Println(isAnagram("rat", "car"))         // false
}
