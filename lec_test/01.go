package main

import (
	"fmt"
	"log"
	"github.com/tarm/serial"
	"strings"
	"syscall"
	"os"
)

func main() {
	Com_Name := "COM3"
	c := &serial.Config{Name: Com_Name, Baud: 115200}
	s, err := serial.OpenPort(c)
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
	//无法串口自动成帧
	var words, word_all string

	file, err := os.OpenFile("./tt.txt", os.O_RDWR | os.O_APPEND | os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	for true {
		buf := make([]byte, 1024)
		n, err := s.Read(buf)
		if err != nil {
				log.Fatal(err)
		}
		words = strings.Trim(fmt.Sprintf("%q", buf[:n]), "\"")
		word_all += words
		word_all = strings.Replace(word_all, "\\r\\n", "\r\n", -1)

		count := strings.Count(word_all, "\r\n")
		if  count > 1 {
			fmt.Printf(word_all)
			// file.write(word_all)
			word_all = ""
		} else if count == 1{
			// index := strings.Index(word_all, "\r\n")
			// fmt.Println(index)
			fmt.Printf(word_all)
			word_all = ""
		}
	}
}