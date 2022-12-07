/*
给定一个字符串数组 ["I","am","stupid","and","weak"]
用 for 循环遍历该数组并修改为["I","am","smart","and","strong"]
*/

package main

import "fmt"

func main() {
	var arr = [5]string{"I", "am", "stupid", "and", "weak"}
	fmt.Println("Before:", arr)

	for i := 0; i < len(arr); i++ {
		if arr[i] == "stupid" {
			arr[i] = "smart"
		}
		if arr[i] == "weak" {
			arr[i] = "strong"
		}
	}
	fmt.Println("After:", arr)
}
