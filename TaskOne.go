package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	//one()
	//two()
	//three()
	//four()
	//five()
	//six()
	//seven()
	eight()
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

func four() {
	//编写一个函数来查找字符串数组中的最长公共前缀。
	//
	//如果不存在公共前缀，返回空字符串 ""。

	strs := []string{"flower", "flow", "floight"}
	prefix := ""
	for i := 0; i < len(strs); i++ {
		flag := true
		for a, v := range strs {
			if a == 0 {
				prefix += runeAt(v, i)
			}
			if runeAt(v, i) != runeAt(prefix, i) {
				prefix = prefix[:i]
				flag = false
			}
		}
		if !flag {
			break
		}
	}
	fmt.Println(prefix)
}

func five() {
	//给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
	//
	//将大整数加 1，并返回结果的数字数组。
	digits := []int{1, 2, 3, 1}
	ndigits := []int{}
	for _, d := range digits {
		ndigits = append(ndigits, d)
	}
	if ndigits[len(digits)-1] == 9 {
		ndigits[len(digits)] = 0
	} else {
		ndigits[len(digits)-1] = ndigits[len(digits)-1] + 1
	}
	fmt.Println(ndigits)
}

func six() {
	//给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
	//不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。可以使用双指针法，
	//一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
	nums := []int{1, 2, 3, 3, 4, 5, 5, 6}
	i := 0
	for j, n := range nums {
		if j == 0 {
			continue
		}
		if nums[i] != n {
			i++
			nums[i] = nums[j]
		}
	}
	fmt.Println(i + 1)
}

func seven() {
	//以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
	//请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
	//可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
	//将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
	intervals := [][]int{{1, 3}, {2, 4}, {8, 10}, {15, 18}}
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})
	nintervals := [][]int{}
	a := 0
	for i, v := range intervals {
		if i == 0 {
			nintervals = append(nintervals, v)
			continue
		}
		if nintervals[a][1] >= v[0] {
			nintervals[a][1] = max(nintervals[a][1], v[1])
		} else {
			a++
			nintervals = append(nintervals, v)
		}
	}
	fmt.Println(nintervals)
}

func eight() {
	// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
	nums := []int{2, 7, 11, 15, 20}
	target := 22
	m := make(map[int]int)
	for _, v := range nums {
		cz := target - v
		m[v] = cz
	}
	var a int
	for k, v := range m {
		if Contains(nums, v) && k != v {
			a = k
			break
		}
	}
	fmt.Println(a, m[a])
}

func Contains[T comparable](nums []T, target T) bool {
	for _, v := range nums {
		if v == target {
			return true
		}
	}
	return false
}

func runeAt(s string, i int) string {
	rs := []rune(s)
	if i < 0 || i >= len(rs) {
		return ""
	}
	return string(rs[i])
}
