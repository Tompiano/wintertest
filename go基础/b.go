package main

import "fmt"

//已知以上代码输出2 1,请说明为什么3未输出
//原因：因为a的值为true，进入了if语句后return退出执行，所以不再打印3

func main() {
	var a = true
	defer func() {
		fmt.Println("1")
	}()
	if a {
		fmt.Println("2")
		return
	}
	defer func() {
		fmt.Println("3")
	}()
}
