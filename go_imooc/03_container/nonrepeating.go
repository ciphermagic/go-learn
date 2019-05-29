// 寻找最长不含有重复字符的子串
// abcabcbb -> abc
// bbbbb -> b
// pwwkew -> wke
package main

import "fmt"

func lengthOfNonRepertingSubStr(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepertingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepertingSubStr("bbbbb"))
	fmt.Println(lengthOfNonRepertingSubStr("pwwkew"))
	fmt.Println(lengthOfNonRepertingSubStr(""))
	fmt.Println(lengthOfNonRepertingSubStr("b"))
	fmt.Println(lengthOfNonRepertingSubStr("abcdef"))
}
