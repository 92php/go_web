package main

import "fmt"

func main() {

	var (
		age     int     = 28
		name    string  = "shirdon"
		balance float32 = 999999.99
	)
	fmt.Println("age = ", age)         //age =  28
	fmt.Println("name = ", name)       //name =  shirdon
	fmt.Println("balance = ", balance) //balance =  1e+06

	var (
		age1     int
		name1    string
		balance1 float32
	)

	fmt.Println("age1 = ", age1)         //age1 =  0
	fmt.Println("name1 = ", name1)       //name1 =
	fmt.Println("balance1 = ", balance1) //balance1 =  0

	var language1 string = "Go"
	var language2 = "Go"
	language3 := "Go"

	fmt.Println("language1 = ", language1)
	fmt.Println("language2 = ", language2)
	fmt.Println("language3 = ", language3)

	var age3, name3, balance3 = 30, "liao", 99.99
	fmt.Println("age3 = ", age3)
	fmt.Println("name3 = ", name3)
	fmt.Println("balance3 = ", balance3)

	age2, name2, balance2 := 30, "shirdon", 999999.99

	d, c := "D", "C"
	c, d = d, c

	println(c)
	println(d)

	println(age2)
	println(name2)
	println(balance2)

}
