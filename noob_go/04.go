package _04

import "unsafe"
import "fmt"


func main() {
	//const同c++
	//只有const初始化的时候可以赋值
	//const格式1：
	//const 变量名1, 变量名2 数据类型
	const _int1 int = 10
	println(_int1)
	//cosnt格式2：(自动获取变量类型)
	const (
		_string1 = "zoulee"
		_int2 = len(_string1)
		_int3 = unsafe.Sizeof(_string1)
	)
	println(_string1, _int2, _int3)
	//枚举1
	//iota默认第一个为0
	//cosnt格式2：(自动获取变量类型)
	const (
		_int4 = iota
		_int5
		_int6
	)
	println(_int4, _int5, _int6)
	//iota特殊例子1
	const (
		a = iota   //0
		b          //1
		c          //2
		d = "ha"   //独立值，iota += 1
		e          //"ha"   iota += 1
		f = 100    //iota +=1
		g          //100  iota +=1
		h = iota   //7,恢复计数
		i          //8
	)
	fmt.Println(a,b,c,d,e,f,g,h,i)
	//iota特殊例子1
	const (
		tt1=1<<iota	//左移 0 位，不变仍为 1。
		tt2=3<<iota	//左移 1 位，变为二进制 110，即 6。
		tt3			//左移 2 位，变为二进制 1100，即 12。
		tt4			//左移 3 位，变为二进制 11000，即 24。
	)
	fmt.Println("tt1=",tt1)
    fmt.Println("tt2=",tt2)
    fmt.Println("tt3=",tt3)
    fmt.Println("tt4=",tt4)
}