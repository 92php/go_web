package main

import "fmt"

var name string = "go"

func myfunc() string {
	defer func() {
		name = "python"
	}()

	fmt.Printf("myfunc 函数里的name：%s\n", name)
	return name
}

func deferCall(){
	defer func1()
	defer func2()
	defer func3()
}

func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}

func change(a, b int) (x, y int) {
	x = a + 100
	y = b + 100

	return   //101, 102
	//return x, y  //101, 102
	//return y, x  //102, 101
}

func main() {
	myname := myfunc()
	fmt.Printf("main 函数里的name: %s\n", name)
	fmt.Println("main 函数里的myname: ", myname)

	//myfunc 函数里的name：go
	//main 函数里的name: python
	//main 函数里的myname:  go

	deferCall() //C B A

	name1 := "go"
	defer fmt.Println(name1)
	name1 = "python"
	defer fmt.Println(name1)
	name1 = "java"
	fmt.Println(name1)
	// java python go

	a := 1
	b := 2
	c, d := change(a, b)
	println(c, d) //101 102

}
