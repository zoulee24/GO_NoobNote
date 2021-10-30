package main

import (
	"github.com/tarm/serial"

	"strings"
	"syscall"
	"strconv"
	// "os"
	"errors"
	"fmt"
)

/**
* @param return
*
* []int		//可用端口编号数组
* int		//可用端口数量
*
**/

func Check_Serial() ([]int, int) {
	Serial_Name := "COM"

	var Serial_exists [50]int
	Serial_count := 0

	for i := 1; i <= 50; i++ {
		
		Com_Name := (Serial_Name + strconv.Itoa(i) )
		c := &serial.Config{Name: Com_Name, Baud: 115200}
		
		if SerialCom, err := serial.OpenPort(c); err != nil {
		} else {
			fmt.Printf("COM%d", i)
			fmt.Printf("存在\r\n")
			Serial_exists[Serial_count] = i
			Serial_count++
			SerialCom.Close()
		}
	}
	Serial_exists_output := Serial_exists[:Serial_count] 
	return Serial_exists_output, Serial_count
}

func (this *Serial_Windows)Serial_Loop(){
	var err error
	var SerialPort *serial.Port
	SerialPort = nil
	for true {
		if this.Serial_state {
			SerialPort, err = this.Open_Serial(4, 115200)
			this.judgeError(err)
			err = this.ReadAShowSerial(SerialPort)
			this.judgeError(err)
		} else {
			if SerialPort != nil {
				SerialPort.Close()
				SerialPort = nil
			}
		}
	}
}

func (this *Serial_Windows)Open_Serial(Serial_index, Baud int)  (*serial.Port, error) {
	Serial_name := "COM"+strconv.Itoa(Serial_index)
	temp := &serial.Config{Name: Serial_name, Baud: 115200}
	Serial_com, err := serial.OpenPort(temp)
	if err != nil {
		if err == syscall.ERROR_FILE_NOT_FOUND {
			err = errors.New("无法找到端口: " + Serial_name)
		} else if err == syscall.ERROR_ACCESS_DENIED {
			err = errors.New("端口: " + Serial_name + " 被占用")
		} else {
			err = errors.New("其他错误：" + err.Error())
		}
		return nil, err
	}
	return Serial_com, err
}

func (this *Serial_Windows)ReadAShowSerial(Serial_com *serial.Port) error{
	//串口读取错误
	var SerialReadError error

	if Serial_com != nil {
		//临时存储变量
		var words, word_all string
		//收到的字节数
		var Rx_num int

		for this.Serial_state {
			buf := make([]byte, 1024)
			Rx_num, SerialReadError = Serial_com.Read(buf)
			if SerialReadError != nil {
				break
			}
			words = strings.Trim(fmt.Sprintf("%q", buf[:Rx_num]), "\"")
			word_all += words
			word_all = strings.Replace(word_all, "\\r\\n", "\r\n", -1)

			count := strings.Count(word_all, "\r\n")
			if  count > 1 {
				this.Show_Text_On_Gui(word_all)
				// fmt.Printf(word_all)
				word_all = ""
			} else if count == 1{
				this.Show_Text_On_Gui(word_all)
				// fmt.Printf(word_all)
				word_all = ""
			}
		}
	} else {
		SerialReadError = errors.New("端口打开错误")
	}
	return SerialReadError
}

func (this *Serial_Windows)judgeError (err error) {
	if err != nil {
		this.Show_Msg_Box("错误", err.Error())
	}
}

/*
func read(){
	Com_Name := "COM4"
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
*/