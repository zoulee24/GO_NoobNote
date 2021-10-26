package _06

import "fmt"

func main() {
	count := 0
	//for循环格式1
    for true {
		if count > 10 {
			break
		}
		//报错syntax error: unexpected ++, expecting comma or )
		//fmt.Println(count++)
		fmt.Println(count)
		count++
    }
	fmt.Println("end1\n")
	count = 0
	//for循环格式2
	for count < 10 {
		fmt.Println(count)
		count++
	}
	fmt.Println("end2\n")
	count = 0
	//for循环格式3
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("end3\n")
	//for循环格式4
	for ; count < 10;  {
		fmt.Println(count)
		count++
	}
	fmt.Println("end4\n")
	//for循环格式5
	numbers := [6]int{1, 2, 3, 4, 5}
	for key, val := range numbers {
		fmt.Print("key: %d val: %d\n", key, val)
	}
	fmt.Println("end5\n")
}