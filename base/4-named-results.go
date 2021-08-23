package main

import "fmt"

// Go 的返回值可以被命名，这种情况可以被视为在函数的顶部的变量

// 什么是返回值可以被命名，就是返回值可以定义出名称来，直接在函数中复制，然后空return以下就可以返回。

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
