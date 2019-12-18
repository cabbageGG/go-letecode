package main

import "fmt"

func longestPalindrome(s string) string {
	// 动态规划法
	// 1 、先找所有长度为1的回文
	// 2、在找所有长度为2的回文
	// 3、再找所有长度为3的回文
	// 依次递推p(i,i+3) = p(i+1, i+2) and s[i]==s[i+3],即长度为l的可由长度为l-2的计算出来
	//  p(i,j) = p(i+1, j-1) and s[i] == s[j]

	// m := len(s)
	// dp := [1000][1000]bool{}
	// var max string
	// for i:=0; i < m; i++{
	//     dp[i][i] = true
	//     max = s[i:i+1]
	// }
	// for i:=0; i < m-1; i++ {
	//     if s[i]==s[i+1]{
	//         dp[i][i+1] = true
	//         max = s[i:i+2]
	//     }
	// }

	// for interval := 2; interval < m; interval++{
	//     for i:=0; i + interval < m; i++{
	//         j := i + interval
	//         if dp[i+1][j-1] && s[i]==s[j]{
	//             dp[i][j] = true
	//             max = s[i:j+1]
	//         }
	//     }
	// }
	// return max
	// 优化：只使用两个数组
	m := len(s)
	dp1 := [1000]bool{} //保存基数长度
	dp2 := [1000]bool{} // 偶数
	var max string
	for i := 0; i < m; i++ {
		dp1[i] = true
		max = s[i : i+1]
	}
	for i := 0; i < m-1; i++ {
		if s[i] == s[i+1] {
			dp2[i] = true
			max = s[i : i+2]
		}
	}
	for interval := 2; interval < m; interval++ {
		for i := 0; i+interval < m; i++ {
			j := i + interval
			if interval%2 == 0 {
				if dp1[i+1] && s[i] == s[j] {
					dp1[i] = true
					max = s[i : j+1]
				}else{
                    dp1[i] = false
                }
			} else {
				if dp2[i+1] && s[i] == s[j] {
					dp2[i] = true
					max = s[i : j+1]
				}else{
                    dp2[i] = false
                }
			}
		}
	}
	return max
}

func main() {
	s := "aaaa"
	r := longestPalindrome(s)
	fmt.Println(r)
}
