package main

import (
	"fmt"
	// "string"

	// "github.com/lxn/walk"
	// . "github.com/lxn/walk/declarative"
)

func main() {
	_, Serial_len := Check_Serial()
	fmt.Printf("端口数量 = %d\r\n", Serial_len)
	
	sw_run, sw := GUI_INIT()

	go sw.Serial_Loop()

	if _, err := sw_run.Run(); err != nil {
		fmt.Println(err)
	}
}
