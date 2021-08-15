package main

import "fmt"

func main() {
	i, j := 42, 2701

	// & 操作符会生成一个指向其操作数的指针
	p := &i
	fmt.Println(p)  // 查看内存地址
	fmt.Println(*p) // 通过指针读取值
	*p = 21
	fmt.Println(i) // 查看值是否改变

	p = &j
	*p = *p / 37 // 通过指针对j进行运算
	fmt.Println(j)

}
