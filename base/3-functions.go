package main

// 函数可以无参，也可以多个参数

import "fmt"

// add 接受两个int类型的参数，返回一个int类型
func add(x int, y int) int {
	return x + y
}

// 函数的参数两个或者多个相同类型的时候，可以缩写，如下
func otherAdd(x, y, z int) int {
	return x + y + z
}

// 函数可以返回多个值，如下
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(otherAdd(1, 2, 3))
	fmt.Println(swap("a", "b"))
}
