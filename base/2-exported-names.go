package main

// 在Go中，如果一个名字以大写字母开头， 那么就是已导出的。

import "fmt"
import "math"

func main() {
	fmt.Println(math.Pi)
	// 任何未导出的名字在该包中都无法访问，如下就会报错
	// fmt.Println(math.pi)
	// out:
	//# command-line-arguments
	//./2-exported-names.go:11:14: cannot refer to unexported Name math.pi
}
