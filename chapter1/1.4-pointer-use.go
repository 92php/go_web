package main

import (
	"fmt"
)


// 交换函数
func exchange(c, d *int) {
	t := *c // 取c指针的值, 赋给临时变量t
	*c = *d // 取d指针的值, 赋给c指针指向的变量
	*d = t // 将c指针的值赋给d指针指向的变量
}


func exchange2(c, d *int) {
	d, c = c, d
}

func main() {
	var score int = 100
	var name string = "Barry"
	// 使用 fmt.Printf 的动词%p输出 score 和 name 变量取地址后的指针值
	fmt.Printf("%p %p", &score, &name)


	var address = "Chengdu,China"         // 准备一个字符串类型
	ptr := &address                       // 对字符串取地址, ptr类型为*string
	fmt.Printf("ptr type: %T\n", ptr)     // 打印ptr的类型   ptr type: *string
	fmt.Printf("address: %p\n", ptr)      // 打印ptr的指针地址 address: 0xc00003c200
	value := *ptr                         // 对指针进行取值操作
	fmt.Printf("value type: %T\n", value) // 取值后的类型 value type: string
	fmt.Printf("value: %s\n", value)      // 指针取值后就是指向变量的值 value: Chengdu,China


	a, b := 6, 8      // 准备两个变量, 赋值6和8
	exchange(&a, &b)  // 交换变量值
	fmt.Println(a, b) // 输出变量值 8 6


	x, y := 6, 8
	exchange2(&x, &y)
	fmt.Println(x, y) //6 8

	var bb int = 66 //定义一个普通类型
	var pb *int = &bb //定义一个指针类型
	println(bb) //66
	println(pb) //0xc00007fe98
	println(*pb) //66
}
