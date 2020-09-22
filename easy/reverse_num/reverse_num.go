package main

import (
	"fmt"
)

//  取模%10 获取末尾最后一个数字, 然后再把num 进行/10 递增然后再取模重复
func main() {
	v := reverse(-2147483412)

	fmt.Printf("%d\n", v)

}

func reverse(x int) int {

	var t int

	for x != 0 {

		// 初始t=0 + 假设x=123 取模为3
		t = t*10 + x%10

		// x=123 取余/10  =12
		x = x / 10

		//越界检查
		if t > (1<<31-1) || t < (-1<<31) {
			return 0
		}
	}
	return t
}
