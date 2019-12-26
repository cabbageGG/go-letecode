package main

func isPowerOfTwo(n int) bool {
	// num := 1
	// for n >=num {
	//     if n == num {
	//         return true
	//     }
	//     num = num << 1
	// }
	// return false
	return n > 0 && (n&(n-1) == 0)
}

func main() {

}
