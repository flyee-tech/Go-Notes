package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// 更灵活的创建数组方式
	primers := []int{1, 2, 3, 4, 5}
	fmt.Println(primers)

	// 可以使用切片来切分数组（包前不包后）
	fmt.Println(primers[1:3])

	// 注意一个点，切换的数据只是创建了一个引用
	a1 := primers[1:3]
	a1[0] = 1000
	fmt.Println(a1)
	fmt.Println(primers)

	// 切片可以忽略上下界，下界的默认值为0，上界的默认值为 length
	fmt.Println(primers[:])
	fmt.Println(primers[:3])

	// 切片拥有长度和容量
	printSlice(primers[:0])
	printSlice(primers[:])
	printSlice(primers[4:]) // !! 这里容量就不是6了，因为容量是从第一个元素开始数

	// 使用make来创建切片（动态数组）
	printSlice(make([]int, 5))
	printSlice(make([]int, 5, 10))

	// 可以使用append函数为切片添加元素
	primers = append(primers, 5, 6, 7, 8)
	printSlice(primers)

	iterArray(primers)

}

func printSlice(s []int) {
	fmt.Printf("%v len=%d cap=%d \n", s, len(s), cap(s))
}

func iterArray(arr []int) {
	// 使用range关键字可以遍历出切片
	for i, v := range arr {
		fmt.Println(i, v)
	}
	// 可以使用 _ 来忽略值
	// for _, v := range arr {
	// for i, _ := range arr {
	// for i := range arr {
}

// Pic 切片练习答案
func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for i := range result {
		result[i] = make([]uint8, dx)
	}
	return result
}
