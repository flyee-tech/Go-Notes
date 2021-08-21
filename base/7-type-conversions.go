package main

import (
	"fmt"
	"math"
)

// go中的类型转换

// T(v) 把值v转换成T类型
// 不同类型的项之间赋值时需要显式转换，如果没有转换代码，程序编译会报错

func main() {
	x, y := 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, f, z)
}
