package _03

import "fmt"

//声明变量格式3:
/*
var{
	变量名1 数据类型
	变量名2 数据类型
}
*/

var{
	_int3 int
	_string3 string
}

func main() {
	//变量只能声明一次 
	//声明变量格式1:
	//var 变量名1, 变量名2 数据类型
	//(不能声明的同时赋值)
	var _int int
	var _string string
	_int = 8
	_string = "你好"
	fmt.Println(_int)
	fmt.Println(_string)
	//声明变量格式2:
	//变量名1, 变量名2 := 初始化数据(自动获取变量类型)
	_int2, _string2 := 10, "你好2"
	fmt.Println(_int2)
	fmt.Println(_string2)
}