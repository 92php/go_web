package main

//go语言for循环
func main() {
	/*	product := 1
			for i := 1; i < 5; i++ {
				product *= i
			}
			println(product) //1*2*3*4 = 24

			m := 0
			for {
				m++
				if m > 50 {
					break
				}
			}
			println(m) //51

		JumpLoop:
			for j := 0; j < 5; j++ {
				for ii := 0; ii < 5; ii++ {
					if ii > 2 {
						break JumpLoop
					}
					fmt.Println(ii) //0 1 2
				}
			}

			jj := 2
			for ; jj > 0; jj-- {
				fmt.Println(jj) // 2 1
			}*/

	/*	var ij int
	JumpLoop1:
		for ; ; ij++ {
			if ij > 10 {
				println(ij)  //11
				break JumpLoop1
			}
		}*/

	/*		var is int
			for {
				if is > 10 {
					break
				}
				is++
				println(is) // 1 2 3 4 5 6 7 8 9 10 11
			}*/

	var i int
	for i <= 20 {
		println(i) //0到20
		i++
	}

}
