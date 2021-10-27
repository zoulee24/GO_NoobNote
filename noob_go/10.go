package _10

import "fmt"
import "unsafe"

func main()  {
	_int1 := 1
	_int1 = 2
	//指针的用法和c++差不多
	//GO没有强制类型转化和编译器隐式转化
	//所有类型转化必须显示
	var _int_ptr *int = &_int1
	fmt.Println(*_int_ptr)
	println("")
	//数据类型的强制转化
	//格式1：
	//目标数据类型(数据)
	var _int2 int = int(2.0)
	println(_int2)
	println("")
	//!注意指针不能强制类型转化	必须要用
	//指针数据类型的强制转化
	//格式2：
	//(目标数据类型指针)unsafe.Pointer(某指针)
	var _int32_ptr *int32 = (*int32)(unsafe.Pointer(_int_ptr)) 
	//int类型 != int32类型
	fmt.Println(*_int32_ptr)
}