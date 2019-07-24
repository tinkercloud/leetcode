package main

import (
	"fmt"
	"regexp"
)

// 给你一个有效的 IPv4 地址 address，返回这个 IP 地址的无效化版本。
// 所谓无效化 IP 地址，其实就是用 "[.]" 代替了每个 "."。

// 示例 1：

// 输入：address = "1.1.1.1"
// 输出："1[.]1[.]1[.]1"
// 示例 2：

// 输入：address = "255.100.50.0"
// 输出："255[.]100[.]50[.]0"

func main() {
	ip := "255.100.50.0"

	newIP := FormatIP(ip)
	fmt.Println(newIP)
}

func FormatIP(address string) string {
	// var newAddress []byte
	re := regexp.MustCompile(`\.`)
	newAddress := re.ReplaceAllString(address, "[.]")

	return newAddress
}
