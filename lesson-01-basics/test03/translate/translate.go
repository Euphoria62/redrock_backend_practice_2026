package main

import "fmt"
func square(n int) int {
	return n * n
}

// TODO（清单第 9 步）：把 reference.c 里的
//
//	int square(int n) { return n * n; }
//
// 改写成 Go 的函数。Go 用 func 开头，参数写成「名字 类型」，
// 返回类型跟在参数列表后面。

func main() {
	total := 0

	for i :=1; i <=5; i++ {
		if i%2==0 {
			fmt.Printf("%d even %d\n", i, square(i))
		}else {
            fmt.Printf("%d odd %d\n", i, square(i))
        }
        total = total + square(i)
	}
	fmt.Printf("total %d\n", total)
}
