package main

import "fmt"

func main() {
	var array = []int{3,4,5}
	//array := []int{6, 8, 10} //定义局部变量
	c := min(array)
	fmt.Println("c = ",c) //c =  3

	a, b := compute(6, 8)
	fmt.Println(a, b) //14 48


	/* 定义局部变量 */
	num1 := 6
	num2 := 8
	fmt.Printf("交换前num1的值为 : %d\n", num1)
	fmt.Printf("交换前num2的值为 : %d\n", num2)
	/* 通过调用函数来交换值 */
	exchange(&num1, &num2)
	fmt.Printf("交换后num1的值 : %d\n", num1)
	fmt.Printf("交换后num2的值 : %d\n", num2)
}

//获取整型数组中的最小值
func min(arr []int) (min int) {
	min = arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return
}

func compute(x, y int) (int, int) {
	return x + y, x * y
}

func sum(x, y int) int {
	return x + y
}

/* 定义相互交换值的函数 */
func exchange(x *int, y *int) int {
	var tmp int
	tmp = *x /* 将 *x 值赋给 tmp */
	*x = *y  /* 将 *y 值赋给 *x */
	*y = tmp /* 将 tmp 值赋给 y*/
	return tmp
}
