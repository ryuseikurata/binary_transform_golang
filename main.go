package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("整数を入力してください=>")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()
	var decimal_number, err = strconv.Atoi(text)
	if err != nil {
		log.Fatal(err)
		return
	}
	if decimal_number == 1 || decimal_number == 0 {
		fmt.Printf("2進数は、%d", decimal_number)
		return
	}
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

	fmt.Printf("2進数は、%s", str)
}
func intSliceToStrSlice(i_list []int) []string {
	s := make([]string, len(i_list))
	for n := range i_list {
		s[n] = strconv.Itoa(i_list[n])
	}

	return s
}
