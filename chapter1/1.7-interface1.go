package main

import "fmt"

type stringer interface {
	string() string
}

type tester interface {
	stringer // 嵌入其他接口
	test()
}

type data struct{}

func (*data) test() {
	fmt.Println("go")
}

func (data) string() string {
	return "hello"
}

// message是一个定义了通知类行为的接口
type Message interface {
	sending()
}

// 定义user及user.notify方法
type User struct {
	name  string
	phone string
}

func (u *User) sending() {
	fmt.Printf("Sending user phone to %s<%s>\n", u.name, u.phone)
}

// 定义admin及admin.message方法
type admin struct {
	name  string
	phone string
}

func (a *admin) sending() {
	fmt.Printf("Sending admin phone to %s<%s>\n", a.name, a.phone)
}

// sendMessage接受一个实现了message接口的值 并发送通知
func sendMessage(n Message) {
	n.sending()
}

func main() {

	var d data
	var t tester = &d
	t.test() //go
	println(t.string()) //hello


	var var1, var2 interface{}
	println(var1 == nil, var1 == var2) //true true
	var1, var2 = 66, 88
	println(var1 == var2) //false


	var a interface{} = func(a int) string {
		return fmt.Sprintf("d:%d", a)
	}
	switch b := a.(type) { // 局部变量b是类型转换后的结果
	case nil:
		println("nil")
	case *int:
		println(*b)
	case func(int) string:
		println(b(66))
	case fmt.Stringer:
		fmt.Println(b)
	default:
		println("unknown")
	}
	//d:66



	// 创建一个user值并传给sendMessage
	bill := User{"Barry", "barry@gmail.com"}
	sendMessage(&bill) //Sending user phone to Barry<barry@gmail.com>

	// 创建一个admin值并传给sendMessage
	lisa := admin{"Jim", "jim@gmail.com"}
	sendMessage(&lisa) //Sending admin phone to Jim<jim@gmail.com>


}
