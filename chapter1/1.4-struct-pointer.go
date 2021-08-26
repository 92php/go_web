package main

import "fmt"

//定义一个结构体
type Books struct {
	title   string
	author  string
	subject string
	press   string
}

func main() {
	/*var bookGo Books     // 声明bookGo为Books类型
	var bookPython Books // 声明bookPython为Books类型

	bookGo.title = "Go Web编程实战派从入门到精通"
	bookGo.author = "廖显东"
	bookGo.subject = "Go语言教程"
	bookGo.press = "电子工业出版社"

	bookPython.title = "Python教程xxx"
	bookPython.author = "张三"
	bookPython.subject = "Python语言教程"
	bookPython.press = "xxx出版社"

	printBook(&bookGo)

	printBook(&bookPython)*/

	/*// 创建一个新的结构体
	fmt.Println(Books{"Go Web编程实战派从入门到精通", "廖显东", "Go语言教程", "电子工业出版社"})
	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go Web编程实战派从入门到精通", author: "廖显东", subject: "Go语言教程", press: "电子工业出版社"})
	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go Web编程实战派从入门到精通", author: "廖显东"})*/

	/*var bookGo Books        //声明bookGo为Books类型
	var bookPython Books        //声明bookPython为Books类型
	// bookGo描述
	bookGo.title = "Go Web编程实战派从入门到精通"
	bookGo.author = "廖显东"
	bookGo.subject = "Go语言教程"
	bookGo.press = "电子工业出版社"

	// bookPython描述
	bookPython.title = "Python教程xxx"
	bookPython.author = "张三"
	bookPython.subject = "Python语言教程"
	bookPython.press = "xxx出版社"

	//打印 bookGo 信息
	fmt.Printf( "bookGo title : %s\n", bookGo.title)
	fmt.Printf( "bookGo author : %s\n", bookGo.author)
	fmt.Printf( "bookGo subject : %s\n", bookGo.subject)
	fmt.Printf( "bookGo press : %s\n", bookGo.press)

	//打印 bookPython 信息
	fmt.Printf( "bookPython title : %s\n", bookPython.title)
	fmt.Printf( "bookPython author : %s\n", bookPython.author)
	fmt.Printf( "bookPython subject : %s\n", bookPython.subject)
	fmt.Printf( "bookPython press : %s\n", bookPython.press)*/


	var bookGo Books     // 声明bookGo为Books类型
	var bookPython Books // 声明bookPython为Books类型

	bookGo.title = "Go Web编程实战派从入门到精通"
	bookGo.author = "廖显东"
	bookGo.subject = "Go语言教程"
	bookGo.press = "电子工业出版社"

	bookPython.title = "Python教程xxx"
	bookPython.author = "张三"
	bookPython.subject = "Python语言教程"
	bookPython.press = "xxx出版社"

	printBook1(bookGo)
	printBook1(bookPython)
}

func printBook1(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book press : %s\n", book.press)
}

func printBook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book press : %s\n", book.press)
}
