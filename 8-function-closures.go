package main

import "fmt"

// !! 函数可以是一个闭包
// !! 什么是闭包？ 闭包是一个函数值，他引用了函数之外的变量（函数和变量绑定在一起）
func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

// adder 返回一个闭包，每个闭包都被绑定在其变量sum上
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// 练习答案，斐波那契数列
func fibonacci() func() int {
	prev := 0
	next := 1
	return func() int {
		tmp := prev
		prev, next = next, prev+next
		return tmp
	}
}
