package main

import (
	"fmt"
	"math/cmplx"
)

// go 的基本类型
// 语法点：变量的声明可以"分组"成一个语法块。
var (
	ToBe          = false
	MaxInt uint64 = 1<<64 - 1
	z             = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Println(ToBe)
	fmt.Println(MaxInt)
	fmt.Println(z)

	// 布尔类型
	var b bool
	fmt.Println(b)

	// 字符串类型
	var str string = "str"
	fmt.Println(str)

	// 整型
	var i1 int
	var i2 int8
	var i3 int16
	var i4 int32
	var i5 int64
	fmt.Println(i1, i2, i3, i4, i5)
	// 还有无符号整型
	var ui1 uint = 1
	fmt.Println(ui1)

	// byte 是 unit8 的别名
	var b1 byte
	fmt.Println(b1)

	// rune 表示int32的别名 也表示 一个Unicode 码点
	var r1 rune
	fmt.Println(r1)

	// 浮点型
	var f1 float32
	var f2 float64
	fmt.Println(f1, f2)

	// complex 类型
	//TODO 中文不知道叫什么?科学记数法？
	var c1 complex64
	fmt.Println(c1)

}
