/*Given a string S, find the longest palindromic substring in S
Input:
S = "aaaabbaa"
Output: aabbaa */

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	fmt.Println("_________________________")
	fmt.Println(longestPalindrome(s))

}

func longestPalindrome(s string) string {
	end := 0
	start := 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		len := max(len1, len2)
		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++

	}

	return right - left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
