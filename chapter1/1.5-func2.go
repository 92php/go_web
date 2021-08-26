package main

import (
	"fmt"
	"math"
)

/* 定义结构体 */
type Circle struct {
	radius float64
}

func main() {
	/* 定义局部变量 */
	num1 := 6
	num2 := 8
	fmt.Printf("交换前num1的值为 : %d\n", num1)
	fmt.Printf("交换前num2的值为 : %d\n", num2)
	/* 通过调用函数来交换值 */
	exchange(num1, num2)
	fmt.Printf("交换后num1的值 : %d\n", num1)
	fmt.Printf("交换后num2的值 : %d\n", num2)


	getPow := func(x float64, y float64) float64 { //声明函数变量
		return math.Pow(x, y)
	}
	fmt.Println(getPow(3, 4)) //使用函数 81


	x, y := 6, 8
	defer func(a int) {
		fmt.Println("defer x, y = ", a, y) //y为闭包引用
	}(x)
	x += 10
	y += 100
	fmt.Println(x, y)
	//16 108
	//defer x, y =  6 108


	var c Circle
	c.radius = 6.00
	fmt.Println("圆的周长 = ", c.getCircumference()) //圆的周长 =  37.68


}

/* 定义相互交换值的函数 */
func exchange(x, y int) int {
	var tmp int
	tmp = x /* 将 x 值赋给 tmp */
	x = y   /* 将 y 值赋给 x */
	y = tmp /* 将 tmp 值赋给 y*/
	return tmp
}

func (c Circle) getCircumference() float64 { //该 method 属于 Circle 类型对象中的方法
	return 2 * 3.14 * c.radius //c.radius 即为 Circle 类型对象中的属性
}
