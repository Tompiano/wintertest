package main

import (
	"fmt"
)

// 不用看了。。。没写对。。。。
func main() {
	var n, time, number int
	var egg []int
	fmt.Scan(&n) //输入有几条船
	for i := 1; i <= n; i++ {
		fmt.Scan(&time) //输入时间和蛋的品质?
		egg, number = removeRepeat(egg, n)
		fmt.Println(number)
	}
}
func removeRepeat(egg []int, n int) (newEgg []int, kind int) {
	//通过去除数组里的重复字符从而让蛋的品质的种类不重复
	//每次判断该品质种类跟后面的是否有重复，没有重复就添加进新的数组里面，有重复就去除
	newEgg = make([]int, n)
	kind = 0
	for i := 0; i < len(egg); i++ {
		ifRepeat := false
		for j := i + 1; j < len(egg); j++ {
			if egg[i] == egg[i+1] {
				ifRepeat = true
				break
			}
		}
		if ifRepeat == false {
			newEgg = append(newEgg, egg[i])
			kind++
		}
	}
	return newEgg, kind
}
