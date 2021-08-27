package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var x float64 = 3.4
	//fmt.Println("type:", reflect.TypeOf(x)) //type: float64

	//var x float64 = 6.8
	//fmt.Println("value:", reflect.ValueOf(x)) //Valueof方法会返回一个Value类型的对象 value: 6.8

	/*var x float64 = 6.8
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type()) //type: float64
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64) //kind is float64: true
	fmt.Println("value:", v.Float()) //value: 6.8*/

	/*var x uint8 = 'x'
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())  //type: uint8
	fmt.Println("kind is uint8: ", v.Kind() == reflect.Uint8) // kind is uint8:  true
	x = uint8(v.Uint())
	fmt.Println("x = ",x) //120*/

	type DefinedInt int
	var x DefinedInt = 8
	v := reflect.ValueOf(x)
	fmt.Println(v) //8


	var name string = "Go Web Program"
	v1 := reflect.ValueOf(name)
	fmt.Println("可写性为:", v1.CanSet()) //可写性为: false



	/*var name1 interface{} = "shirdon"
	fmt.Printf("原始接口变量的类型为 %T，值为 %v \n", name1, name1) //原始接口变量的类型为 string，值为 shirdon
	t := reflect.TypeOf(name1)
	v2 := reflect.ValueOf(name1)
	// 从接口变量到反射对象
	fmt.Printf("从接口变量到反射对象：Type对象的类型为 %T \n", t) //从接口变量到反射对象：Type对象的类型为 *reflect.rtype
	fmt.Printf("从接口变量到反射对象：Value对象的类型为 %T \n", v2) //从接口变量到反射对象：Value对象的类型为 reflect.Value
	// 从反射对象到接口变量
	i := v2.Interface()
	fmt.Printf("从反射对象到接口变量：新对象的类型为 %T 值为 %v \n", i, i) //从反射对象到接口变量：新对象的类型为 string 值为 shirdon*/



	/*var name3 string = "Go Web Program"
	fmt.Println("真实 name 的原始值为：", name3) //真实 name 的原始值为： Go Web Program
	v3 := reflect.ValueOf(&name3)
	v4 := v3.Elem()
	v4.SetString("Go Web Program2")
	fmt.Println("通过反射对象进行更新后，真实 name 变为：", name3) //通过反射对象进行更新后，真实 name 变为： Go Web Program2*/



	var name4 string = "Go Web Program"
	v5 := reflect.ValueOf(&name4)
	fmt.Println("v5 可写性为:", v5.CanSet()) //v5 可写性为: false
	v6 := v5.Elem()
	fmt.Println("v6 可写性为:", v6.CanSet()) //v6 可写性为: true


}
