package main

import (
	"fmt"
	"strings"
)

type Vertex2 struct {
	Lat, Long float64
}

var m map[string]Vertex2

func main() {
	// 使用make函数返回给定类型的映射
	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{
		40.1, -72,
	}
	fmt.Println(m)

	// 插入和修改元素
	m["test"] = Vertex2{1, 2}
	m["test"] = Vertex2{2, 3}

	// 获取元素
	test := m["test"]
	fmt.Println(test)

	// 删除元素
	delete(m, "test")

	// 检测某个值是否存在
	test, ok := m["test"]
	fmt.Println(test, ok)

	WordCount("hello world")

}

// WordCount 映射练习答案
func WordCount(s string) map[string]int {
	result := make(map[string]int)
	wordArr := strings.Fields(s)
	for _, v := range wordArr {
		_, ok := result[v]
		if ok {
			result[v] = result[v] + 1
		} else {
			result[v] = 1
		}
	}
	return result
}
