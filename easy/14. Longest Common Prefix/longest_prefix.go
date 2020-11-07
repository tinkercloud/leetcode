package main

import "fmt"

func main() {
	strarr := []string{"a"}
	res := longestCommonPrefix(strarr)
	fmt.Println(res)
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	pref := strs[0]

	for i := 0; i < len(pref); i++ {

		// 读取字符串数组第一个字符串
		for j := 1; j < len(strs); j++ {
			fmt.Printf("array[i]:%d\tstrs[j]:%d\tstrs[j][i]:%d\tpref[i]:%d\n", i, len(strs[j]), strs[j][i], pref[i])
			// 最小的字符串大小匹配     数组的第j个字符串的第i个字符 和 pref第i个字符不相等
			if i == len(strs[j]) || strs[j][i] != pref[i] {
				return pref[:i]
			}

		}

	}
	return pref
}
