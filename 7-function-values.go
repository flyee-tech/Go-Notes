package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// 函数也是值，和其他变量是一样的
func main() {
	// 创建一个函数的变量值
	hypot := func(x, y float64) float64 {
		return math.Max(x, y)
	}
	// 可以调用函数
	fmt.Println(hypot(5, 12))
	// 可以当成参数传入其他函数中
	fmt.Println(compute(hypot))
	// 函数可以传入符合参数的函数
	fmt.Println(compute(math.Max))
}
