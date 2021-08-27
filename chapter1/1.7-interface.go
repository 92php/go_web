package main

import "fmt"

type Num int

func (x Num) Equal(i Num) bool {
	return x == i
}

func (x Num) LessThan(i Num) bool {
	return x < i
}

func (x Num) BiggerThan(i Num) bool {
	return x > i
}

func (x *Num) Multiple(i Num) {
	*x = *x * i
}

func (x *Num) Divide(i Num) {
	*x = *x / i
}

type NumI interface {
	Equal(i Num) bool
	LessThan(i Num) bool
	BiggerThan(i Num) bool
	Multiple(i Num)
	Divide(i Num)
}

type NumI2 interface {
	Equal(i Num) bool
	LessThan(i Num) bool
	MoreThan(i Num) bool
}

func Len(array interface{}) int {
	var length int //数组的长度
	if array == nil {
		length = 0
	}
	switch array.(type) {
	case []int:
		length = len(array.([]int))
	case []string:
		length = len(array.([]string))
	case []float32:
		length = len(array.([]float32))

	default:
		length = 0
	}
	fmt.Println(length)

	return length
}

type InterfaceFile interface {
	Read()
	Write()
}

type InterfaceReader interface {
	Read()
}

type File struct {
}

func (f *File) Read() {

}

func (f *File) Write() {

}

//接口1
type Interface1 interface {
	Write(p []byte) (n int, err error)
}

//接口2
type Interface2 interface {
	Close() error
}

//接口组合
type InterfaceCombine interface {
	Interface1
	Interface2
}

func main() {

	slice := make([]int, 0)
	slice = append(slice, 6, 7, 8)
	var I interface{} = slice
	if res, ok := I.([]int); ok {
		fmt.Println(res) //[6 7 8]
		fmt.Println(ok) //true
	}


	f := new(File)

	var f1 InterfaceFile = f
	var f2 InterfaceReader = f1

	var f3 InterfaceReader = new(File)

	// 接口查询
	if f5, ok := f3.(InterfaceFile); ok {
		fmt.Println(f5) //&{}
	}

	// 询问接口它指向的对象是否是某个类型
	if f6, ok := f3.(*File); ok {
		fmt.Println(f6) //&{}
	}

	fmt.Println(f1, f2, f3) //&{} &{} &{}



}
