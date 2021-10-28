package _12

import "fmt"

func main() {
	//数组切片
	//格式1：
	//var 变量名字 []数据类型 = make([]数据类型, 开始, 结束)
	var numbers = make([]int,3,5)
	_int_s := []int{1, 2, 3, 4, 5}
	printSlice(numbers)
	printSlice(_int_s)
	//格式2:
	//变量名字 := 数组变量名字[开始:结束]
	//	(很像python)
	_int_qp := _int_s[2:4]
	printSlice(_int_qp)
}

func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
