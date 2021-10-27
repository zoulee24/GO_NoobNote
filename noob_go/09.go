package _09

import (
	"fmt"
)

func main() {
	//定义数组
	//格式1：
	//var 变量名字 [大小]数据类型
	var _floats32_s1 [5]float32
	for key, _ := range _floats32_s1 {
		_floats32_s1[key] = float32(key)
		fmt.Println(_floats32_s1)
	}
	fmt.Println("")
	// //格式2：
	// //var 变量名字 = [大小]数据类型{数据1, 数据2}
	// var _int32_s = [3]int32{10, 1, 2} 
	//格式3：
	//变量名字 := [大小]数据类型{数据1, 数据2}
	_floats32_s2 := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	fmt.Println(_floats32_s2)
	fmt.Println("\n")
	//格式4：
	//变量名字 := [...]数据类型{数据1, 数据2}
	//...代表自动获取数组大小
	_floats32_s3 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(_floats32_s3)
}