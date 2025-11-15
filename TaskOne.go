package main

import (
	"fmt"
	"strconv"
)

func main() {
	//one()
	//two()
	three()
}

func one() {
	//给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
	//找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map
	//数据结构来解决，例如通过 map
	//记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
	arr := []int{4, 1, 2, 1, 2, 3, 6, 6, 3}
	ctMap := make(map[int]int)
	for _, v := range arr {
		ctMap[v]++
	}
	fmt.Println(ctMap)
	for k, v := range ctMap {
		if v == 1 {
			fmt.Println(k)
		}
	}
}

func two() {
	flag := true
	//判断一个整数是否是回文数
	num := 12345678987654321
	numstr := strconv.Itoa(num)
	i := 0
	for j := len(numstr) - 1; j >= 0; j-- {
		if numstr[i] != numstr[j] {
			flag = false
		}
		if i == j {
			break
		}
		i++
	}
	fmt.Println(flag)
}

func three() {
	//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
	//
	//有效字符串需满足：
	//
	//左括号必须用相同类型的右括号闭合。
	//左括号必须以正确的顺序闭合。
	//每个右括号都有一个对应的相同类型的左括号。
	str := "(()){]{}[]"
	m := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	stack := []string{}
	a := 0
	for _, v := range str {
		if len(stack) == 0 || stack[a-1] != m[string(v)] {
			stack = append(stack, string(v))
			a++
		}
		if stack[a-1] == m[string(v)] {
			stack = append(stack[:a-1], stack[a:]...)
			a--
		}
	}
	fmt.Println(stack)
	fmt.Println(len(stack) == 0)
}
