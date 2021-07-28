package main

import (
	"fmt"
	"strings"
)

func main() {
	const decimal_number = 50
	const binary = 2
	var result_list []int

	var prev_n = decimal_number
	var n = decimal_number / binary
	for n != 1 {
		var result = prev_n - n*binary
		result_list = append(result_list, result)
		prev_n = n
		n = n / binary
	}
	// 1になった時の処理
	var last_value = prev_n - n*binary
	result_list = append(result_list, last_value)

	// 1を最後に追加
	result_list = append(result_list, 1)

	// 逆順にする
	for i := 0; i < len(result_list)/2; i++ {
		result_list[i], result_list[len(result_list)-i-1] = result_list[len(result_list)-i-1], result_list[i]
	}

	var r = intSliceToStrSlice(result_list)

	var str = strings.Join(r, "")

	fmt.Printf(str)
}
func intSliceToStrSlice(i_list []int) []string {
	s := make([]string, len(i_list))
	for n := range i_list {
		s[n] = string(rune(i_list[n]))
		fmt.Printf(s[n])
	}

	return s
}
