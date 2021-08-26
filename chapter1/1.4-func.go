package main

import "fmt"

func main() {
	// 定义匿名函数并赋值给f变量
	f := func(data int) {
		fmt.Println("hi, this is a closure", data)
	}
	// 此时f变量的类型是func(), 可以直接调用
	f(6)

	//直接声明并调用
	func(data int) {
		fmt.Println("hi, this is a closure, directly", data)
	}(8)


	sli := []int{1, 6, 8}
	// 使用匿名函数打印切片的内容
	visitPrint(sli, func(value int) {
		fmt.Println(value)  // 1 6 8
	})


}


// 遍历切片中每个元素，通过给定的函数访问元素
func visitPrint(list []int, f func(int)) {
	for _, value := range list {
		f(value)
	}
}
