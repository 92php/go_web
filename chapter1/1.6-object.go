package main

import (
	"fmt"
)

// 正方形
type Square struct {
	sideLen float32
}

// 三角形结构体
type Triangle struct {
	Bottom float32
	Height float32
}


type Student struct {
	Name  string
	score float32
}

// 计算三角形面积
func (t *Triangle) Area() float32 {
	return (t.Bottom * t.Height) / 2
}

// 接口 Shape
type Shape interface {
	Area() float32
}

// 计算正方形的面积
func (sq *Square) Area() float32 {
	return sq.sideLen * sq.sideLen
}

type Engine interface {
	Start()
	Stop()
}

type Bus struct {
	Engine // 包含 Engine 类型的匿名字段
}

func (c *Bus) Working() {
	c.Start() //开动汽车
	c.Stop()  //停车
}

func main() {
	/*r := Triangle{6, 8}
	// 调用 Area() 方法，计算面积
	fmt.Println(r.Area()) //24

	student := new(Student)
	student.Name = "shirdon"
	student.score = 9
	println(student.Name) //shirdon
	println(student.score) //+9.000000e+000*/


	/*s := new(person.Student)
	s.SetName("shirdon")
	fmt.Println(s.GetName())*/


	t := &Triangle{6, 8}
	s := &Square{8}
	shapes := []Shape{t, s}    // 创建一个 Shape 类型的数组
	for n, _ := range shapes { // 迭代数组上的每一个元素并调用 Area() 方法
		fmt.Println("图形数据: ", shapes[n])
		fmt.Println("它的面积是: ", shapes[n].Area())
	}
    //图形数据:  &{6 8}
    //它的面积是:  24
    //图形数据:  &{8}
    //它的面积是:  64


}
