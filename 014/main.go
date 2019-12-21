package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	commonStr := getCommonPrefix(strs[0], strs[1])
	for j := 2; j < len(strs); j++ {
		temp := getCommonPrefix(commonStr, strs[j])
		if len(temp) < len(commonStr) {
			commonStr = temp
		}
	}
	return commonStr
}

func getCommonPrefix(str1, str2 string) string {
	i := 0
	for i < len(str1) && i < len(str2) {
		if str1[i] != str2[i] {
			break
		}
		i++
	}
	return str1[:i]
}

func main() {

	a := []string{"qqw", "qww"}
	fmt.Println(longestCommonPrefix(a))

}
