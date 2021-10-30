package main

import (
	"fmt"
	"github.com/tarm/serial"
	"syscall"
	"os"

	"log"
	"strings"

	// "time"
	
	"github.com/lxn/walk"
)

type SerialData struct {
	index 	uint64 //下标
	data	float64 //
}

func Serial(mw *MyMainWindow) {
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

	var All_Index uint64
	All_Index = 0

	
	// ok:=false
	// for ok != true {
	// 	if mw.serialcanvas != nil {
	// 		sss, err := walk.NewCosmeticPen(walk.PenSolid, walk.RGB(0, 0, 0))
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 		defer sss.Dispose()
	// 		p1, p2 := walk.Point{X: 150, Y: 456}, walk.Point{X: 245, Y: 247}
	// 		err = mw.serialcanvas.DrawLinePixels(sss, p1, p2)
	// 		p1, p2 = walk.Point{X: 150, Y: 100}, walk.Point{X: 400, Y: 400}
	// 		err = mw.serialcanvas.DrawLinePixels(sss, p1, p2)
	// 		// time.Sleep(time.Second * 5)
	// 		// mw.paintWidget.SetPaintMode(0)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 			ok = false
	// 		} else {
	// 			ok = true
	// 		}
	// 	} else {
	// 		fmt.Println("serialcanvas is nil")
	// 	}
	// }

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
			word_all_s := strings.Split(word_all,"\r\n")
			for _, _word := range word_all_s {
				index_start := strings.Index(_word, "=")
				if index_start != -1 {
					fmt.Println(_word[index_start+1:])
					All_Index++
				} else {
					fmt.Println(_word)
				}
			}
			word_all = ""
		} else if count == 1{
			
			index_start := strings.Index(word_all, "=")
			index_end := strings.Index(word_all, "\r")
			fmt.Println(word_all[index_start+1:index_end])
			
			word_all = ""
			All_Index++
		}
	}
}