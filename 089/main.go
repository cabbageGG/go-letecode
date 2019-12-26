package main

import "fmt"

// 格雷编码是一个二进制数字系统，在该系统中，两个连续的数值仅有一个位数的差异。

// 给定一个代表编码总位数的非负整数 n，打印其格雷编码序列。格雷编码序列必须以 0 开头。

// G(n+1) = (0+G(n)) + (1+'G(n)) // 'G(n)是G(n)的逆序

func grayCode(n int) []int {
	res, head := []int{0}, 1
	for i := 1; i <= n; i++ {
		L := len(res)
		for j := L - 1; j >= 0; j-- {
			res = append(res, res[j]+head)
		}
		head = head * 2
	}
	return res
}

func main() {
	n := 2
	fmt.Println(grayCode(n))
}
