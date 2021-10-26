package _07

import "fmt"

func main() {
	_int1 := 2
	//switch格式1
	switch (_int1) {
		case 0:
			fmt.Println("0")
		case 2:
			fmt.Println("2")
		default:
			fmt.Println("default")
	}
	fmt.Println("")
	//switch格式2
	//fallthrough(强制执行下一个语句，很像c++没加break)
	switch {
		case _int1 == 0:
			fmt.Println("0")
		case _int1 == 2:
			fmt.Println("2")
			fallthrough
		default:
			fmt.Println("default")
	}
	fmt.Println("")
	//switch格式3
	//通过判断数据类型
	var _int2 interface {}
	switch _type := _int2.(type) {
		case nil:
			fmt.Println("类型:%T", _type)
		case int:
			fmt.Println("int")
		case func(int) float64:
			fmt.Println("func(int)")
		case bool, string:
			fmt.Println("bool or string")
		case float64:
			fmt.Println("float64")
	}
	fmt.Println("")
}
