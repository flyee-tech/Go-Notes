package main

import "fmt"

// var 语句可以定义变量，变量可以出现在包和函数级别

var c, java, python bool

// 声明变量的时候，可以包含【初始值】

var i, j int = 1, 2

// 在包含初始值的情况下可以【省略类型】

var aa, bb, cc = true, false, "yes"

// 函数外的语句都必须以关键字开头，所以短变量不能在函数外使用
// 如下会报错
// jj := 1

func main() {
	// 从这里可以看到，函数内的变量优先级大于函数外的变量
	var i int
	fmt.Println(i, c, java, python)
	fmt.Println(i, j, aa, bb, cc)

	// 短变量声明
	k := 3
	fmt.Println(k)
}
