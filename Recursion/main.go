package main

import "fmt"

// Base condition is length of curr == n*2
// Open always lessthen n
// Close always  lessthen open
func generateParenthesis(n int) []string {
	var res []string

	var backtrack func(curr string, open, close int)

	backtrack = func(curr string, open, close int) {
		if len(curr) == n*2 {
			res = append(res, curr)
			return
		}

		if open < n {
			backtrack(curr+"(", open+1, close)
		}

		if close < open {
			backtrack(curr+")", open, close+1)
		}
	}

	backtrack("", 0, 0)

	return res
}

func towerOfHanoi(n int, src, heper, dest string) {
	if n == 1 {
		fmt.Printf("Move desk 1 src %s to -> %s\n", src, dest)
		return
	}
	towerOfHanoi(n-1, src, dest, heper)
	fmt.Printf("Move desk %d from %s to ->  %s\n", n, src, dest)
	towerOfHanoi(n-1, heper, src, dest)
}

func main() {
	// problem :: Generate Parentheses - (Leecode  22)
	// fmt.Println(generateParenthesis(3))

	// problem :: Letter Combination of Phone Numbers - (Leecode  17)

	// problem :: Tower of Hanoi
	// https://www.youtube.com/watch?v=3aag8Wpv8XM
	n := 4
	towerOfHanoi(n, "A", "B", "C")
}
