package _08

import "fmt"

func main() {
	c, d := 10, 20
	var res int

	//例1
	res = Max_1(c, d)
	fmt.Printf("res1=%d\n\n", res )

	//例2
	c, d = 100, 20
	res = Max_2(c, d)
	fmt.Printf("res2=%d\n\n", res )

	//例3
	c, d = 5, 2
	fmt.Printf("swap前c=%d\td=%d\n", c, d )
	c, d = swap(c, d)
	fmt.Printf("swap后c=%d\td=%d\n\n", c, d )

	//例4
	show(1, "hello")
}
//可以有多个return

//格式1
// func 函数名(变量名1, 变量名2 输入值数据类型) 返回值数据类型

func Max_1(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
//Max_1 功能上等于 Max_2
func Max_2(a, b int) int {
	var rest int
	if a < b {
		rest = b
	} else {
		rest = a
	}
	return rest
}
//可以返回多个值，和pyhton很像

//格式2
// func 函数名(变量名1, 变量名2 输入值数据类型) (返回值数据类型1, 返回值数据类型2)

func swap(a, b int) (int, int) {
	if a > b{
		temp := b
		b = a
		a = temp
	}
	return a, b
}

//格式2
// func 函数名(变量名1, 变量名2 输入值数据类型, 变量名3 输入值数据类型)
func show(a int, b string) {
	fmt.Println(a, b)
}