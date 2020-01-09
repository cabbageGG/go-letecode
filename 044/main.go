package main

import "fmt"

func isMatch(s string, p string) bool {
	i := 0
	j := 0

	strStarIdx := 0  // string中匹配'*'的位置
	patStarIdx := -1 // pattern中出现'*'的位置
	for i < len(s) {
		if j < len(p) && (s[i] == p[j] || p[j] == '?') {
			i++
			j++
		} else if j < len(p) && p[j] == '*' {
			strStarIdx = i
			// 这里记录下pattern中'*'出现的位置，是为了可以找到pattern中'*'后面的那个字符
			// 因为在后面str匹配不上的时候，要认为str中的字符被pattern中的'*'匹配了，继续往后比较
			patStarIdx = j
			// 这里没有i++，因为*可以匹配0个字符
			j++
		} else if patStarIdx != -1 {
			// 如果str和pattern的当前字符没有匹配上，并且pattern中的前一个字符是'*'，
			// 那就让'*'尽可能多的匹配，都认为是和'*'匹配和(直到在str中找到匹配的字符)

			j = patStarIdx + 1 // j指向pattern中'*'后面的那个字符，每次不匹配后，都从'*'后的字符重新匹配 // 很关键，这里就是回溯！！！
			strStarIdx += 1    // 如果str不匹配pattern中'*'后面的那个字符，则让str往后走，都认为是和'*'匹配
			i = strStarIdx
		} else {
			return false
		}
	}

	for ; j < len(p); j++ {
		if p[j] != '*' {
			return false
		}
	}

	return true
}

func main() {
	s := "cdddadab"
	p := "*ab"
	fmt.Println(isMatch(s, p))
}
