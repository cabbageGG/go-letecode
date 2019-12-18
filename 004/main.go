package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	L := len(nums1) + len(nums2)
	resIndex := L / 2
	n, i, j, v1, v2 := 0, 0, 0, 0, 0
	for n <= resIndex {
		v1 = v2
		if i >= len(nums1) {
			v2 = nums2[j]
			j++
		} else if j >= len(nums2) {
			v2 = nums1[i]
			i++
		} else {
			if nums1[i] < nums2[j] {
				v2 = nums1[i]
				i++
			} else {
				v2 = nums2[j]
				j++
			}
		}
		n++
	}
	if L%2 == 1 {
		return float64(v2)
	} else {
		return float64(v1+v2) / 2
	}
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
