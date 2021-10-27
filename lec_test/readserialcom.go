package readserialcom

import (
	"syscall"
	"fmt"
	"os"
	// "errors"
)

func main() {
	Com_Name := "COM4"
	filename := "\\\\.\\" + Com_Name
	fmt.Printf("filename = %s\n", Com_Name)
	h, err := syscall.CreateFile(
		syscall.StringToUTF16Ptr(filename),
		syscall.GENERIC_READ | syscall.GENERIC_WRITE,
		0, 
		nil, 
		syscall.OPEN_EXISTING, 
		syscall.FILE_ATTRIBUTE_NORMAL | syscall.FILE_FLAG_OVERLAPPED,
		0)
	//如果有错误就进入
	if err != nil {
		if err == syscall.ERROR_FILE_NOT_FOUND {
			fmt.Println( "无法找到端口: " + Com_Name )
		} else if err == syscall.ERROR_ACCESS_DENIED {
			fmt.Println( "端口: " + Com_Name + " 被占用" )
		} else {
			fmt.Print( "其他错误: ")
			fmt.Println( err )
		}
	}
	f := os.NewFile(uintptr(h), filename)
	defer func() {
		if err != nil {
			f.Close()
		}
	}()
}