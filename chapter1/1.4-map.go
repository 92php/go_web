package main

import "fmt"

func main() {
	var literalMap map[string]string
	var assignedMap map[string]string
	literalMap = map[string]string{"first": "go", "second": "web"}
	createdMap := make(map[string]float32)
	assignedMap = literalMap
	createdMap["k1"] = 99
	createdMap["k2"] = 199
	assignedMap["second"] = "program"
	fmt.Printf("Map literal at \"first\" is: %s\n", literalMap["first"])    //Map literal at "first" is: go
	fmt.Printf("Map created at \"k2\" is: %f\n", createdMap["k2"])          //Map created at "k2" is: 199.000000
	fmt.Printf("Map assigned at \"second\" is: %s\n", literalMap["second"]) //Map assigned at "second" is: program
	fmt.Printf("Map literal at \"third\" is: %s\n", literalMap["third"])    //Map literal at "third" is:

	achievement := map[string]float32{
		"zhangsan": 99.5, "lisi": 88,
		"wangwu": 96, "lidong": 100,
	}
	fmt.Println("achievement = ", achievement) //achievement =  map[lidong:100 lisi:88 wangwu:96 zhangsan:99.5]

	createdMap1 := new(map[string]float32)
	fmt.Println("createdMap1 = ", createdMap1) //createdMap1 =  &map[]

	map2 := make(map[string]float32, 100)
	fmt.Println("map2 = ", map2) //map2 =  map[]

}
