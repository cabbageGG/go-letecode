package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	var res []int
	for len(matrix) > 0 {

		res = append(res, matrix[0]...)
		matrix = matrix[1:]

		for i := 0; i < len(matrix); i++ {
			temp := matrix[i]
			if len(temp) > 0 {
				res = append(res, temp[len(temp)-1])
				matrix[i] = temp[:len(temp)-1]
			}
		}

		if len(matrix) > 0 {
			temp := matrix[len(matrix)-1]
			if len(temp) > 0 {
				for i := len(temp) - 1; i >= 0; i-- {
					res = append(res, temp[i])
				}
			}
			matrix = matrix[:len(matrix)-1]
		}

		for i := len(matrix) - 1; i >= 0; i-- {
			temp := matrix[i]
			if len(temp) > 0 {
				res = append(res, temp[0])
				matrix[i] = temp[1:]
			}
		}
	}
	return res
}

func main() {
	// var m [][]int
	// a := []int{1, 2, 3}
	// b := []int{4, 5, 6}
	// c := []int{7, 8, 9}
	// m = append(m, a, b, c)
	n := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	fmt.Println(spiralOrder(n))
}
