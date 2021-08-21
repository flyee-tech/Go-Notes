package main

import "fmt"

// 数值常量 是 高精度的值

const (
	Big = 1 << 100
	// Small 其实就是 1 << 1 二进制就是10 也就是2
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	//如下会报错，数值超过int的最大值
	//fmt.Println(needInt(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
