//  input 121
// output: true

// input -121

// output: false

package main

import "fmt"

func main() {
	n := 100
	fmt.Println(isPalindrome(n))
}

// 这里更优化解法就是只需要取一半比较就可以了,但要注意奇数和偶数问题
// 暂时想不到更好的解法
// 从社区获取到了解法, 只需要比较一半就可以了
func isPalindrome(x int) bool {

	//
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	var t int

	for t < x {
		t = t*10 + x%10
		x /= 10
	}

	return x == t || t/10 == x

	// if x >= 0 {
	// 	var t, d int
	// 	d = x
	// 	for x != 0 {
	// 		t = t*10 + x%10
	// 		x = x / 10
	// 	}
	// 	return d == t
	// }
	// return false

}
