package main

import "fmt"

// 声明全局变量
var global int

// 声明全局变量
var global2 int = 8

func main() {

	/*// 声明局部变量
	var local1, local2 int

	// 初始化参数
	local1 = 8
	local2 = 10
	global = local1 * local2

	fmt.Printf("local1 = %d, local2 = %d , global = %d\n", local1, local2, global)*/

	// 声明局部变量
	var local1, local2, local3 int

	//初始化参数
	local1 = 8
	local2 = 10
	local3 = local1 * local2

	fmt.Printf(" local1 = %d, local2 = %d , local3 = %d\n", local1, local2, local3) //local1 = 8, local2 = 10 , local3 = 80


	// 声明局部变量
	var global2 int = 999

	fmt.Printf("global2 = %d\n", global2) //global2 = 999
}
